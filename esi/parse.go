package esi

import (
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"os"
	"strconv"
	"strings"
)

func pickName(names []*ESILangName) string {
	for _, n := range names {
		if n.LanguageID == "1033" {
			return n.Name
		}
	}

	if len(names) != 0 {
		return strings.TrimSpace(names[0].Name)
	} else {
		return ""
	}
}

func ParseESIFile(filename string, emitPDOs, emitObjects bool) ([]*ESIDevice, error) {
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

	if len(esiData.VendorNames) != 0 {
		esiData.VendorName = esiData.VendorNames[0].Name
		for _, v := range esiData.VendorNames {
			if v.LanguageID == "1033" { // English
				esiData.VendorName = strings.TrimSpace(v.Name)
			}
		}
	}

	groupMap := make(map[string]string)

	// Beckhoff's ESI groups have names in DE and EN; we want the English names.
	for _, g := range esiData.Groups {
		for _, n := range g.Names {
			if n.LanguageID == "1033" { // English
				g.EnglishName = strings.TrimSpace(n.Name)
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
				d.EnglishName = strings.TrimSpace(n.Name)
			}
		}
		for _, u := range d.URLs {
			if u.LanguageID == "1033" { // English
				d.EnglishURL = strings.TrimSpace(u.URL)
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
		if !emitPDOs {
			d.TxPDOs = nil
			d.RxPDOs = nil
		}

		for _, txpdo := range d.TxPDOs {
			txpdo.Index = fixHexFormat(txpdo.Index, 4)
			txpdo.PDOName = pickName(txpdo.LangNames)
			for _, entry := range txpdo.Entries {
				entry.PDOName = pickName(entry.LangNames)
				entry.Index = fixHexFormat(entry.Index, 4)
				entry.SubIndex = fixHexFormat(entry.SubIndex, 2)
			}
		}
		for _, rxpdo := range d.RxPDOs {
			rxpdo.Index = fixHexFormat(rxpdo.Index, 4)
			rxpdo.PDOName = pickName(rxpdo.LangNames)
			for _, entry := range rxpdo.Entries {
				entry.PDOName = pickName(entry.LangNames)
				entry.Index = fixHexFormat(entry.Index, 4)
				entry.SubIndex = fixHexFormat(entry.SubIndex, 2)
			}
		}

		// Blank out object data if it's disabled via the command line.
		if !emitObjects {
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
		d.IDs = []*ESIDeviceID{}

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
