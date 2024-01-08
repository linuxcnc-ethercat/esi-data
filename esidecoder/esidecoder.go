package main

import (
	"flag"
	"fmt"
	"github.com/linuxcnc-ethercat/esi-data/esi"
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

func rewriteBeckhoffNumbers(devices []*esi.ESIDevice) {
	for _, d := range devices {
		typesplit := strings.Split(d.ShortType, "-")
		if strings.HasPrefix(d.Vendor, "Beckhoff") && len(typesplit) > 2 {
			// We're going to assume that this is a device
			// like a Beckhoff AXS5101-0000-0010, which
			// should really be a AXS5101-0000 revision
			// 0010.
			lastpart := typesplit[len(typesplit)-1] // Fetch the last chunk of the name.
			rev, err := strconv.ParseInt(d.RevisionNo, 0, 64)
			if err != nil {
				fmt.Printf("*** unable to parse Beckoff revision number, continuing: %v\n", err)
				continue
			}

			// Turn the top 4 digits of the revision number into a 0-padded decimal number.
			decimalRev := fmt.Sprintf("%04d", rev>>16)

			if lastpart == decimalRev {
				newType := strings.Join(typesplit[0:len(typesplit)-1], "-")
				d.EnglishName = strings.ReplaceAll(d.EnglishName, d.ShortType, newType)
				d.ShortType = newType
				
			}

		}
	}
}

func main() {
	flag.Parse()

	devices := []*esi.ESIDevice{}

	files, err := os.ReadDir(*esiDirFlag)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".xml") {
			fmt.Fprintf(os.Stderr, "Parsing %s... ", file.Name())

			d, err := esi.ParseESIFile(filepath.Join(*esiDirFlag, file.Name()), *pdosFlag, *objectsFlag)
			if err != nil {
				panic(err)
			}

			fmt.Fprintf(os.Stderr, "%d devices\n", len(d))

			devices = append(devices, d...)
		}
	}

	rewriteBeckhoffNumbers(devices)

	devices2, err := esi.MergeDevices(devices)

	y, err := yaml.Marshal(devices2)
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
