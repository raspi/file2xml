package main

import (
	"os"
	"io"
	"fmt"
	"encoding/xml"
	"flag"
	"strings"
	"errors"
)

type Byte struct {
	Position uint64 `xml:"pos,attr"`
	Bits     [8]Bit `xml:"Bit,"`
}

type Bit struct {
	Position uint8 `xml:"pos,attr"`
	Enabled  bool `xml:",attr"`
}

func HasRequiredCommandLineArguments(required []string, seen map[string]bool) (err error) {
	err = nil
	var errs []string
	for _, req := range required {
		if !seen[req] {
			errs = append(errs, fmt.Sprintf("Error: Missing required -%s argument/flag.", req))
		}
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "\n"))
	}

	return err
}

func main() {
	var err error

	sourceFile := flag.String("f", "", "File name to read (data.dat)")
	outputFile := flag.String("o", "", "Output XML file name to write (out.xml)")

	flag.Parse()

	required := []string{"f", "o"}

	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) {
		seen[f.Name] = true
	})

	err = HasRequiredCommandLineArguments(required, seen)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}

	_, err = os.Stat(*sourceFile)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "File not found: '%v'", *sourceFile)
			os.Exit(2)
		} else {
			panic(err)
		}

	}

	// Source file
	f, err := os.Open(*sourceFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Output XML file
	of, err := os.Create(*outputFile)
	if err != nil {
		panic(err)
	}
	defer of.Close()

	// XML Header
	_, err = of.Write([]byte(xml.Header))
	if err != nil {
		panic(err)
	}

	_, err = of.Write([]byte("<Data>\n"))
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 1024*1024)

	position := uint64(0)

	for {
		readBytes, err := f.Read(buffer)

		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}

		if readBytes == 0 {
			break
		}

		for idx := 0; idx < readBytes; idx++ {
			var b Byte
			b.Position = position

			// One byte
			data := buffer[idx]

			var bit [8]Bit

			// Read each bit
			bit[7].Position = 7
			bit[7].Enabled = (data & 1) != 0
			bit[6].Position = 6
			bit[6].Enabled = (data & 2) != 0
			bit[5].Position = 5
			bit[5].Enabled = (data & 4) != 0
			bit[4].Position = 4
			bit[4].Enabled = (data & 8) != 0
			bit[3].Position = 3
			bit[3].Enabled = (data & 16) != 0
			bit[2].Position = 2
			bit[2].Enabled = (data & 32) != 0
			bit[1].Position = 1
			bit[1].Enabled = (data & 64) != 0
			bit[0].Position = 0
			bit[0].Enabled = (data & 128) != 0

			b.Bits = bit
			position++

			xmldata, err := xml.MarshalIndent(&b, "  ", "  ")
			if err != nil {
				panic(err)
			}

			of.Write(xmldata)
			of.Write([]byte("\n"))
		}

	}

	_, err = of.Write([]byte("</Data>\n"))
	if err != nil {
		panic(err)
	}


}
