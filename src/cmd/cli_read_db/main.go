package main

import (
	"fmt"
	"os"

	"github.com/zkhrg/go_day01/pkg/dbreader"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "incorrect number of arguments\n")
		return
	}
	if os.Args[1] != "-f" {
		fmt.Fprintf(os.Stderr, "need '-f' at call\n")
		return
	}
	filename := os.Args[2]
	reader, err := dbreader.ReaderByFileExtension(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
	if err := reader.ReadFile(filename); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	if err := reader.Print(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
}
