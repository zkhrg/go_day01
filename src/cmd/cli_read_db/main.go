package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/zkhrg/go_day01/pkg/dbreader"
)

func main() {
	filePtr := flag.String("f", "", "path to file")
	flag.Parse()
	err := runByFileExtension(*filePtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}

func runByFileExtension(filename string) error {
	var err error
	ext := filepath.Ext(filename)
	switch ext {
	case ".json":
		json_db_reader := dbreader.JSONDBReader{}
		json_db_reader.ReadFile(filename)
		json_db_reader.Print()
	case ".xml":
		xml_db_reader := dbreader.XMLDBReader{}
		xml_db_reader.ReadFile(filename)
		xml_db_reader.Print()
	default:
		err = errors.New("file extension is not provided")
	}
	return err
}
