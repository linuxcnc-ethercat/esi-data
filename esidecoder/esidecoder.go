package main

import (
	"flag"
	"fmt"
	"github.com/linuxcnc-ethercat/esi-data/esi"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

var (
	esiDirFlag  = flag.String("esi_directory", "/tmp/esi", "Directory that contains ESI XML files")
	pdosFlag    = flag.Bool("pdos", false, "Include PDO data in output")
	objectsFlag = flag.Bool("objects", false, "Include object data in output")
	output      = flag.String("output", "", "Output file, defaults to stdout")
)

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
