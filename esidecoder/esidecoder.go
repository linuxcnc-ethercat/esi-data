package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	esiDirFlag  = flag.String("esi_directory", "/tmp/esi", "Directory that contains ESI XML files")
	pdosFlag    = flag.Bool("pdos", false, "Include PDO data in output")
	objectsFlag = flag.Bool("objects", false, "Include object data in output")
	output      = flag.String("output", "", "Output file, defaults to stdout")
)

type ESIName struct {
	Name       string `xml:",chardata"`
	LanguageID string `xml:"LcId,attr"`
}

type ESIDeviceURL struct {
	URL        string `xml:",chardata"`
	LanguageID string `xml:"LcId,attr"`
}

type ESIDeviceInfo struct {
	Current         string              `xml:"Electrical>EBusCurrent" yaml:",omitempty"`
	RequestTimeout  int                 `xml:"Mailbox>Timeout>RequestTimeout" yaml:",omitempty"`
	ResponseTimeout int                 `xml:"Mailbox>Timeout>ResponseTimeout" yaml:",omitempty"`
	Ports           []ESIDeviceInfoPort `xml:"Port" yaml:",omitempty"`
}

type ESIDeviceInfoPort struct {
	Type  string
	Label string
}

type ESIProfile struct {
	ProfileNo   int            `xml:"ProfileNo" yaml:",omitempty"`
	Channels    int            `xml:"ChannelCount" yaml:",omitempty"`
	AddInfo     int            `xml:"AddInfo" yaml:",omitempty"` // what is "AddInfo?"
	DataTypes   []*ESIDataType `xml:"Dictionary>DataTypes>DataType"`
	DataTypeMap map[string]*ESIDataType
	Objects     []*ESIObject `xml:"Dictionary>Objects>Object" yaml:"Objects,omitempty"`
}

type ESIDataType struct {
	Name     string                `xml:"Name"`
	BitSize  int                   `xml:"BitSize"`
	SubItems []*ESIDataTypeSubItem `xml:"SubItem"`
}

type ESIDataTypeSubItem struct {
	SubIndex  int    `xml:"SubIdx"`
	Name      string `xml:"Name"`
	Type      string `xml:"Type"`
	BitSize   int    `xml:"BitSize"`
	BitOffset int    `xml:"BitOffs"`
	Access    string `xml:"Flags>Access"`
	Category  string `xml:"Flags>Category"`
}

type ESIObject struct {
	Index    string              `xml:"Index" yaml:",omitempty"`
	Name     string              `xml:"Name" yaml:",omitempty"`
	TypeName string              `xml:"Type" yaml:",omitempty"`
	BitSize  int                 `xml:"BitSize" yaml:",omitempty"`
	SubItems []*ESIObjectSubItem `xml:"Info>SubItem" yaml:",omitempty"`
	// Info > (DefaultData | SubItem | ...
	// Flags > (Access | Category | ...
}

type ESIObjectSubItem struct {
	Name        string `xml:"Name"\ yaml:",omitempty"`
	DefaultData string `xml:"Info>DefaultData" yaml:",omitempty"`
	Type        string `yaml:"Type,omitempty"`
	BitSize     int    `yaml:"BitSize,omitempty"`
	BitOffset   int    `yaml:"BitOffs,omitempty"`
	Access      string `yaml:"Access,omitempty"`
	Category    string `yaml:"Category,omitempty"`
}

type ESIDevice struct {
	//Physics string `xml:"Physics,attr" yaml:",omitempty"`

	Type struct {
		Type            string `xml:",chardata"`
		ProductCode     string `xml:"ProductCode,attr"`
		RevisionNo      string `xml:"RevisionNo,attr"`
		CheckRevisionNo string `xml:"CheckRevisionNo,attr"`
	} `xml:"Type" yaml:"-"`

	Names     []ESIName      `xml:"Name" yaml:",omitempty"`
	URLs      []ESIDeviceURL `xml:"URL" yaml:",omitempty"`
	Info      ESIDeviceInfo  `xml:"Info" yaml:"Info,omitempty"`
	GroupType string         `xml:"GroupType" yaml:",omitempty"` // maps to ESIGroup.Type
	//Fmmu      []string       `xml:"Fmmu" yaml:",omitempty,flow"` // "Inputs"/"Outputs"

	ShortType   string      `yaml:"Type"`
	ProductCode string      `yaml:"ProductCode"`
	RevisionNo  string      `yaml:"RevisionNo"`
	EnglishURL  string      `yaml:"URL"`
	EnglishName string      `yaml:"Name"`
	GroupName   string      `yaml:"DeviceGroup"`
	Vendor      string      `yaml:"Vendor"`
	VendorID    string      `yaml:"VendorID"`
	Profile     ESIProfile  `yaml:"Profile,omitempty"`
	TxPDOs      []*ESIRxPDO `xml:"TxPdo" yaml:"TxPDOs,omitempty"`
	RxPDOs      []*ESITxPDO `xml:"RxPdo" yaml:"RxPDOs,omitempty"`
}

type ESIRxPDO struct {
	Index   string         `xml:"Index" yaml:"Index,omitempty"`
	Name    string         `xml:"Name" yaml:"Name,omitempty"`
	Entries []*ESIPDOEntry `xml:"Entry" yaml:"Entry,omitempty"`
}

type ESITxPDO struct {
	Index   string         `xml:"Index" yaml:"Index,omitempty"`
	Name    string         `xml:"Name" yaml:"Name,omitempty"`
	Entries []*ESIPDOEntry `xml:"Entry" yaml:"Entry,omitempty"`
}

type ESIPDOEntry struct {
	Name     string `xml:"Name" yaml:"Name,omitempty"`
	Index    string `xml:"Index" yaml:"Index,omitempty"`
	SubIndex string `xml:"SubIndex" yaml:"SubIndex,omitempty"`
	BitLen   int    `xml:"BitLen" yaml:"BitLen,omitempty"`
	DataType string `xml:"DataType" yaml:"DataType,omitempty"`
}

type ESIGroup struct {
	Type        string    `xml:"Type"`
	Names       []ESIName `xml:"Name"`
	EnglishName string
}

type ESIData struct {
	EtherCATInfo xml.Name     `xml:"EtherCATInfo"`
	VendorID     string       `xml:"Vendor>Id"`
	VendorName   string       `xml:"Vendor>Name"`
	Groups       []*ESIGroup  `xml:"Descriptions>Groups>Group"`
	Devices      []*ESIDevice `xml:"Descriptions>Devices>Device"`
}

func parseESIFile(filename string) ([]*ESIDevice, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %v", err)
	}

	esiData := &ESIData{}
	devices := []*ESIDevice{}

	decoder := xml.NewDecoder(f)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(esiData)
	if err != nil {
		return nil, fmt.Errorf("unable to parse xml file %q: %v", filename, err)
	}

	groupMap := make(map[string]string)

	// Beckhoff's ESI groups have names in DE and EN; we want the English names.
	for _, g := range esiData.Groups {
		for _, n := range g.Names {
			if n.LanguageID == "1033" { // English
				g.EnglishName = n.Name
				groupMap[g.Type] = g.EnglishName
			}
		}
	}

	// We have a fair amount of post-processing to do per Device.
	//
	// - We want to get the English name and URLs
	// - We want to include the vendor and vendor ID (correctly formatted)
	// - We want to match the group name to the group type provided
	// - We want to reformat the product code and revision to be 0xXXX-format hex numbers instead of #xXXX
	// - We want to strip extra data (non-English names and URLs, group data, etc
	for _, d := range esiData.Devices {
		d.Vendor = esiData.VendorName
		d.VendorID = fixHexFormat(esiData.VendorID, 8)

		for _, n := range d.Names {
			if n.LanguageID == "1033" { // English
				d.EnglishName = n.Name
			}
		}
		for _, u := range d.URLs {
			if u.LanguageID == "1033" { // English
				d.EnglishURL = u.URL
			}
		}

		d.Profile.DataTypeMap = make(map[string]*ESIDataType)
		// Build type map
		for _, t := range d.Profile.DataTypes {
			d.Profile.DataTypeMap[t.Name] = t
		}
		d.Profile.DataTypes = nil

		// Blank out PDO data if it's disabled via the command
		// line.
		if !*pdosFlag {
			d.TxPDOs = nil
			d.RxPDOs = nil
		}
		
		for _, txpdo := range d.TxPDOs {
			txpdo.Index = fixHexFormat(txpdo.Index, 4)
			for _, entry := range txpdo.Entries {
				entry.Index = fixHexFormat(entry.Index, 4)
				entry.SubIndex = fixHexFormat(entry.SubIndex, 2)
			}
		}
		for _, rxpdo := range d.RxPDOs {
			rxpdo.Index = fixHexFormat(rxpdo.Index, 4)
			for _, entry := range rxpdo.Entries {
				entry.Index = fixHexFormat(entry.Index, 4)
				entry.SubIndex = fixHexFormat(entry.SubIndex, 2)
			}
		}

		// Blank out object data if it's disabled via the command line.
		if !*objectsFlag {
			d.Profile.Objects = nil
		}

		// For now, we only want to keep 0x6xxx, 0x7xxx, and 0x8xxx objects.
		// Otherwise the generated ESI file is *huge*.
		newObjects := []*ESIObject{}
		for _, o := range d.Profile.Objects {
			typeObject := d.Profile.DataTypeMap[o.TypeName]
			o.Index = fixHexFormat(o.Index, 4)

			if typeObject != nil {
				submap := make(map[string]*ESIDataTypeSubItem)
				for _, ts := range typeObject.SubItems {
					submap[ts.Name] = ts
				}
				for _, s := range o.SubItems {
					subtype := submap[s.Name]
					if subtype != nil {
						s.Type = subtype.Type
						s.BitSize = subtype.BitSize
						s.BitOffset = subtype.BitOffset
						s.Access = subtype.Access
						s.Category = subtype.Category
					}
				}
			}

			switch o.Index[0:3] {
			case "0x6": // input, keep
				newObjects = append(newObjects, o)
			case "0x7": // output, keep
				newObjects = append(newObjects, o)
			default: // other, delete for now

			}
		}
		d.Profile.Objects = newObjects
		d.Profile.DataTypeMap = nil // Don't want to marshal it.

		d.Type.ProductCode = fixHexFormat(d.Type.ProductCode, 8)
		d.Type.RevisionNo = fixHexFormat(d.Type.RevisionNo, 8)
		d.GroupName = groupMap[d.GroupType]
		d.ShortType = d.Type.Type
		d.ProductCode = d.Type.ProductCode
		d.RevisionNo = d.Type.RevisionNo
		d.GroupType = ""
		d.Names = nil
		d.URLs = nil

		devices = append(devices, d)
	}
	return devices, nil
}

func fixHexFormat(in string, length int) string {
	// No need to "fix" empty strings
	if len(in) == 0 {
		return in
	}

	// XML likes #xXXXX, change to 0xXXX
	if in[0] == '#' {
		out := []rune(in)
		out[0] = '0'
		v, err := strconv.ParseInt(string(out), 0, 64)
		if err != nil {
			panic(err)
		}
		return fmt.Sprintf("0x%0*x", length, v)
	}

	// Okay, it's probably just an integer.  Beckhoff likes that.  Convert to hex.
	i, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("0x%0*x", length, i)
}

func main() {
	flag.Parse()

	devices := []*ESIDevice{}

	files, err := os.ReadDir(*esiDirFlag)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".xml") {
			fmt.Fprintf(os.Stderr, "Parsing %s... ", file.Name())

			d, err := parseESIFile(filepath.Join(*esiDirFlag, file.Name()))
			if err != nil {
				panic(err)
			}

			fmt.Fprintf(os.Stderr, "%d devices\n", len(d))

			devices = append(devices, d...)
		}
	}

	y, err := yaml.Marshal(devices)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stderr, "Writing %d bytes to output file\n", len(y))

	if *output != "" {
		err = os.WriteFile(*output, y, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println(string(y))
	}
}
