package main

import (
	"flag"
	"fmt"
	"github.com/linuxcnc-ethercat/esi-data/esi"
	"gopkg.in/yaml.v3"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

var (
	esiFileFlag   = flag.String("esi_file", "esi.yml", "File to read from")
	outputDirFlag = flag.String("output_dir", "devices", "Output directory")
)

func formatRevname(rev string) string {
	// If the revision is a hex number like 0x00020000, with the
	// lower 16 bits clear, then drop the lower 16 bits, convert
	// to deciman, and return as 'rX', where X is the decimal
	// version of the top 16 bits of the revision number *minus
	// 16*.  Else return the original string.
	if len(rev) < 2 {
		return fmt.Sprintf("rev %s", rev)
	}

	if rev[0:2] == "0x" {
		r, err := strconv.ParseInt(rev, 0, 64)
		if err == nil { //&& (r&0xffff == 0) {
			return fmt.Sprintf("r%d", r>>16-16)
		}
	}
	return fmt.Sprintf("rev %s", rev)
}

func main() {
	flag.Parse()

	devices := []*esi.ESIDevice{}

	file, err := os.Open(*esiFileFlag)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Reading %s...\n", *esiFileFlag)

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&devices)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Found %d entries\n", len(devices))

	devs := map[string]map[string]*esi.ESIDevice{}
	for _, d := range devices {
		for _, id := range d.IDs {
			devtype := id.Type
			devrev := id.RevisionNo
			if devs[devtype] == nil {
				devs[devtype] = make(map[string]*esi.ESIDevice)
			}
			devs[devtype][devrev] = d
		}
	}

	for d, revs := range devs {
		fmt.Printf("==> %s\n", d)
		if len(strings.Split(d, "-")) > 2 {
			// skip  foo-bar-baz devices for now
			continue
		}

		for r, rev := range revs {
			fmt.Printf("    %s = %s, %s\n", r, rev.IDs[0].Type, rev.IDs[0].RevisionNo)
		}
		f, err := os.Create(filepath.Join(*outputDirFlag, fmt.Sprintf("%s.md", url.QueryEscape(d))))
		if err != nil {
			panic(err)
		}
		defer f.Close()

		createPageFor(f, d, revs)
	}
}

func sortRevs(revs map[string]*esi.ESIDevice) []string {
	r := []string{}
	for k := range revs {
		r = append(r, k)
	}
	slices.Sort(r)
	return r
}

func getIDFor(ids []*esi.ESIDeviceID, t, r string) *esi.ESIDeviceID {
	for _, i := range ids {
		if i.Type == t && i.RevisionNo == r {
			return i
		}
	}
	return nil // Really shouldn't ever happen
}

func checkColspan(row []string, column int) int {
	span := 0
	val := row[column]
	for i := column; i < len(row); i++ {
		if row[i] == val {
			span++
		} else {
			return span
		}
	}
	return span
}

func printTableRow(f io.Writer, row []string, class string, align string) {
	if class != "" {
		class = fmt.Sprintf("class=%q", class)
	}

	// This is ugly; it'd be better to fix this before we get
	// here, but we want empty cells to actually stay empty, not
	// be <pre></pre>
	for i:=0;i<len(row); i++ {
		if row[i] == "<pre></pre>" {
			row[i]=""
		}
	}
	
	fmt.Fprintf(f, "<tr %s>\n", class)
	for i := 0; i < len(row); i++ {
		column := row[i]
		colspan := checkColspan(row, i)
		class := ""

		if i==0 {
			class = "class=\"first\""
		}

		if colspan > 1 {
			fmt.Fprintf(f, "<td %s colspan=%d align=%q>%s</td>\n", class, colspan, align, column)
		} else {
			fmt.Fprintf(f, "<td %s>%s</td>\n", class,column)
		}
		i += colspan - 1
	}
	fmt.Fprintf(f, "</tr>\n")
}

func printTableRowSpan(f io.Writer, row []string, class string, rowspan int) {
	if class != "" {
		class = fmt.Sprintf("class=%q", class)
	}
	fmt.Fprintf(f, "<tr %s>\n", class)
	fmt.Fprintf(f, "<td class=\"first\" rowspan=%d valign=top>%s</td>\n", rowspan, row[0])

	for i := 1; i < len(row); i++ {
		column := row[i]
		colspan := checkColspan(row, i)
		if colspan > 1 {
			fmt.Fprintf(f, "<td colspan=%d align=\"left\">%s</td>\n", colspan, column)
		} else {
			fmt.Fprintf(f, "<td>%s</td>\n", column)
		}
		i += colspan - 1
	}
	fmt.Fprintf(f, "</tr>\n")
}

func formatEquivDevices(device *esi.ESIDevice, devname string) string {
	devs := []string{}
	for _, id := range device.IDs {
		if id.Type != devname {
			devs = append(devs, fmt.Sprintf("<a href=\"%s\">%s %s</a>", url.QueryEscape(id.Type), id.Type, formatRevname(id.RevisionNo)))
		}
	}

	slices.Sort(devs)
	return strings.Join(devs, "<br/>")
}

type pdoline struct {
	key    string
	output string
}

// Format Tx PDOs.  We want to be able to merge across revisions and
// *only* show changes, so we need to break each line up on its own,
// and then we'll reformat them later once we've matched them across
// all revs.  So, the format is basically <key>:<value>, where <key>
// is more or less sorted in order.  Then we'll iterate through N
// parallel arrays (below) and pull the lowest remaining key.
func formatTxPDOs(device *esi.ESIDevice) []pdoline {
	lines := []pdoline{}

	for _, pdo := range device.TxPDOs {
		lines = append(lines, pdoline{key: pdo.Index, output: fmt.Sprintf("%s: %s", pdo.Index, pdo.Name)})

		for _, entry := range pdo.Entries {
			index := entry.Index

			if index[0:3] != "0x6" {
				continue // Don't show 0x0000 gap entries or entries that are outside of the TX space
			}

			subindex := entry.SubIndex
			if len(subindex) > 2 {
				subindex = subindex[2:] // strip leading "0x"
			}
			lines = append(lines, pdoline{key: fmt.Sprintf("%s %s:%s", pdo.Index, index, subindex), output: fmt.Sprintf("  %s:%s  %-30s  %s", index, subindex, entry.Name, entry.DataType)})
		}
	}

	return lines
}

// Format Rx PDOs.  We want to be able to merge across revisions and
// *only* show changes, so we need to break each line up on its own,
// and then we'll reformat them later once we've matched them across
// all revs.  So, the format is basically <key>:<value>, where <key>
// is more or less sorted in order.  Then we'll iterate through N
// parallel arrays (below) and pull the lowest remaining key.
func formatRxPDOs(device *esi.ESIDevice) []pdoline {
	lines := []pdoline{}

	for _, pdo := range device.RxPDOs {
		lines = append(lines, pdoline{key: pdo.Index, output: fmt.Sprintf("%s: %s", pdo.Index, pdo.Name)})

		for _, entry := range pdo.Entries {
			index := entry.Index

			if index[0:3] != "0x7" {
				continue // Don't show 0x0000 gap entries or entries that are outside of the RX space
			}

			subindex := entry.SubIndex
			if len(subindex) > 2 {
				subindex = subindex[2:] // strip leading "0x"
			}
			lines = append(lines, pdoline{key: fmt.Sprintf("%s %s:%s", pdo.Index, index, subindex), output: fmt.Sprintf("  %s:%s  %-30s  %s", index, subindex, entry.Name, entry.DataType)})
		}
	}

	return lines
}

func mergePDOLines(pdolines [][]pdoline) [][]string {
	columns := len(pdolines)
	currentline := make([]int, columns)
	lastline := make([]int, columns)
	results := [][]string{}

	for i := range pdolines {
		lastline[i] = len(pdolines[i])
	}

	// Look at the current line across all columns and pick the
	// lowest `key`.  Then, for each column that has that key as
	// the `currentline` in pdolines, emit the output into
	// `results`.  Then move `currentline` forward for each column
	// that emitted something.
	for {
		// Check to see if we've used all of our input
		done := true
		for i := range currentline {
			if currentline[i] < lastline[i] {
				done = false
			}
		}

		if done {
			return results
		}

		output := make([]string, columns)

		// Find the lowest key across all of the current lines
		lowestKey := ""
		for i := range currentline {
			if currentline[i] < lastline[i] {
				l := pdolines[i][currentline[i]].key
				if lowestKey == "" || l < lowestKey {
					lowestKey = l
				}
			}
		}

		// Emit into `output` if the current line equals `lowestkey`, and move `currentline` forward.
		for i := range currentline {
			//output[i]="-"
			if currentline[i] < lastline[i] && pdolines[i][currentline[i]].key == lowestKey {
				output[i] = pdolines[i][currentline[i]].output
				currentline[i]++
			}
		}
		results = append(results, output)
	}
}

func createPageFor(f io.Writer, devname string, revs map[string]*esi.ESIDevice) error {
	sortedRevs := sortRevs(revs)
	revIDs := map[string]*esi.ESIDeviceID{}
	columns := len(sortedRevs) + 1

	for _, r := range sortedRevs {
		fmt.Printf("**** looking for %q in %v\n", r, revs)
		revIDs[r] = getIDFor(revs[r].IDs, devname, r)
		if revIDs[r] == nil {
			return fmt.Errorf("getIDFor(.., %q, %q)==nil", devname, r)
		}

	}

	name := revIDs[sortedRevs[0]].Name
	vendor := revIDs[sortedRevs[0]].Vendor
	url := revIDs[sortedRevs[0]].URL
	brand := strings.Split(vendor, " ")[0]

	fmt.Fprintf(f, "# %s %s\n", brand, devname)
	fmt.Fprintf(f, "\n")
	fmt.Fprintf(f, "%s\n\n", name)
	fmt.Fprintf(f, "%s\n\n", vendor)
	fmt.Fprintf(f, "Documentation: <a href=%q>%s</a>\n\n", url, url)

	fmt.Fprintf(f, "## Revisions and PDOs\n")

	fmt.Fprintf(f, "<table>\n")

	div := "<pre>"
	enddiv := "</pre>"

	row := make([]string, columns)
	row[0] = "Revision"
	// Revision name header line
	for c, r := range sortedRevs {
		row[c+1] = div+formatRevname(r)+enddiv
	}
	printTableRow(f, row, "", "center")

	row = make([]string, columns)
	row[0] = "Name"
	for c, r := range sortedRevs {
		row[c+1] = div+revIDs[r].Name+enddiv
	}
	printTableRow(f, row, "", "center")

	row = make([]string, columns)
	row[0] = "PID"
	for c, r := range sortedRevs {
		row[c+1] = div+revIDs[r].ProductCode+enddiv
	}
	printTableRow(f, row, "", "center")

	row = make([]string, columns)
	row[0] = "Revision No"
	for c, r := range sortedRevs {
		row[c+1] = div+revIDs[r].RevisionNo+enddiv
	}
	printTableRow(f, row, "", "center")

	row = make([]string, columns)
	row[0] = "Same PDOs as"
	for c, r := range sortedRevs {
		row[c+1] = div+formatEquivDevices(revs[r], devname)+enddiv
	}
	printTableRow(f, row, "", "center")

	row = make([]string, columns)
	row[0] = "TxPDOs"
	txpdodata := [][]pdoline{}

	for _, r := range sortedRevs {
		txpdodata = append(txpdodata, formatTxPDOs(revs[r]))
	}

	txlines := mergePDOLines(txpdodata)

	if len(txlines) > 0 {
		class := "txpdo"
		txline := 0
		fmt.Printf("** Merged, has %d lines (columns=%d)\n", len(txlines), columns)

		row = make([]string, columns+1)
		row[0] = "TX PDOs"
		for c := 0; c < columns-1; c++ {
			if txlines[txline][c] != "" {
				row[c+1] = fmt.Sprintf("<pre>%s</pre>", txlines[txline][c])
				if string(txlines[txline][c][0]) != " " {
					class = "txpdo pdosection"
				}
			} else {
				row[c+1] = ""
			}
		}
		printTableRowSpan(f, row, class, len(txlines))
		txline++

		for {
			class := "txpdo"
			if txline >= len(txlines) {
				break
			}

			row := make([]string, columns-1)
			for c := 0; c < columns-1; c++ {
				if txlines[txline][c] != "" {
					row[c] = fmt.Sprintf("<pre>%s</pre>", txlines[txline][c])
					if string(txlines[txline][c][0]) != " " {
						class = "txpdo pdosection"
					}
				} else {
					row[c] = ""
				}
			}
			printTableRow(f, row, class, "left")
			txline++
		}
	}

	

	row = make([]string, columns)
	row[0] = "RxPDOs"
	rxpdodata := [][]pdoline{}

	for _, r := range sortedRevs {
		rxpdodata = append(rxpdodata, formatRxPDOs(revs[r]))
	}

	rxlines := mergePDOLines(rxpdodata)

	if len(rxlines) > 0 {
		class := "rxpdo"
		rxline := 0
		fmt.Printf("** Merged, has %d lines (columns=%d)\n", len(rxlines), columns)

		row = make([]string, columns+1)
		row[0] = "RX PDOs"
		for c := 0; c < columns-1; c++ {
			if rxlines[rxline][c] != "" {
				row[c+1] = fmt.Sprintf("<pre>%s</pre>", rxlines[rxline][c])
				if string(rxlines[rxline][c][0]) != " " {
					class = "rxpdo pdosection"
				}
			} else {
				row[c+1] = ""
			}
		}
		printTableRowSpan(f, row, class, len(rxlines))
		rxline++

		for {
			class := "rxpdo"
			if rxline >= len(rxlines) {
				break
			}

			row := make([]string, columns-1)
			for c := 0; c < columns-1; c++ {
				if rxlines[rxline][c] != "" {
					row[c] = fmt.Sprintf("<pre>%s</pre>", rxlines[rxline][c])
					if string(rxlines[rxline][c][0]) != " " {
						class = "rxpdo pdosection"
					}
				} else {
					row[c] = ""
				}
			}
			printTableRow(f, row, class, "left")
			rxline++
		}
	}


	fmt.Fprintf(f, "</table>\n")

	return nil
}
