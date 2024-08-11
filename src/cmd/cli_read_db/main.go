package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zkhrg/go_day01/pkg/dbreader"
	"github.com/zkhrg/go_day01/pkg/flaghelper"
)

func main() {
	filename := flag.String("f", "", "file to read (json will be printed xml and vice versa)")
	flag.Parse()

	if err := flaghelper.CheckLengthFlags(os.Args, []string{"-f"}); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}

	reader, err := dbreader.ReaderByFileExtension(*filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}
	if err := reader.ReadFile(*filename); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	if err := reader.Print(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
}
