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
	for i:=column;i<len(row);i++ {
		if row[i]==val {
			span++
		} else {
			return span
		}
	}
	return span
}

func printTableRow(f io.Writer, row []string) {
	fmt.Fprintf(f, "<tr>\n")
	for i:=0;i<len(row);i++ {
		column := row[i]
		colspan := checkColspan(row, i)
		if colspan>1 {
			fmt.Fprintf(f, "<td colspan=%d align=\"center\">%s</td>\n", colspan, column)
		} else {
			fmt.Fprintf(f, "<td>%s</td>\n", column)
		}
		i += colspan-1
	}
	fmt.Fprintf(f, "</tr>\n")
}

func formatEquivDevices(device *esi.ESIDevice, devname string) string {
	devs := []string{}
	for _, id := range device.IDs {
		if id.Type != devname {
			devs = append(devs, fmt.Sprintf("<a href=\"%s.md\">%s %s</a>", url.QueryEscape(id.Type), id.Type, formatRevname(id.RevisionNo)))
		}
	}

	slices.Sort(devs)
	return strings.Join(devs, "<br/>")
}

func createPageFor(f io.Writer, devname string, revs map[string]*esi.ESIDevice) error {
	sortedRevs := sortRevs(revs)
	revIDs := map[string]*esi.ESIDeviceID{}
	columns := len(sortedRevs)+1

	for _, r := range sortedRevs {
		fmt.Printf("**** looking for %q in %v\n", r, revs)
		revIDs[r] = getIDFor(revs[r].IDs, devname, r)
		if revIDs[r] == nil {
			return fmt.Errorf("getIDFor(.., %q, %q)==nil", devname, r)
		}
	}
	fmt.Fprintf(f, "# %s\n", devname)
	fmt.Fprintf(f, "\n")
	fmt.Fprintf(f, "## Revisions\n")

	fmt.Fprintf(f, "<table>\n")

	row := make([]string, columns)
	row[0]=""
	// Revision name header line
	for c, r := range sortedRevs {
		row[c+1] = formatRevname(r)
	}
	printTableRow(f, row)

	row = make([]string, columns)
	row[0]="Name"
	for c, r := range sortedRevs {
		row[c+1] = revIDs[r].Name
	}
	printTableRow(f, row)

	row = make([]string, columns)
	row[0]="PID"
	for c, r := range sortedRevs {
		row[c+1] = revIDs[r].ProductCode
	}
	printTableRow(f, row)

	row = make([]string, columns)
	row[0]="Revision No"
	for c, r := range sortedRevs {
		row[c+1] = revIDs[r].RevisionNo
	}
	printTableRow(f, row)

	row = make([]string, columns)
	row[0]="Same PDOs as"
	for c, r := range sortedRevs {
		row[c+1] = formatEquivDevices(revs[r],devname)
	}
	printTableRow(f, row)

	
	fmt.Fprintf(f, "</table>\n")

	return nil
}
