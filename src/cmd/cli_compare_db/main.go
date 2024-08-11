package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zkhrg/go_day01/pkg/dbcomparator"
	"github.com/zkhrg/go_day01/pkg/flaghelper"
)

func main() {
	old := flag.String("old", "", "old version db")
	new := flag.String("new", "", "new version db")

	flag.Parse()

	if err := flaghelper.CheckLengthFlags(os.Args, []string{"--new", "--old"}); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}

	if *old == "" {
		fmt.Println("Error: --old is required")
		flag.Usage()
		os.Exit(1)
	}

	if *new == "" {
		fmt.Println("Error: --new is required")
		flag.Usage()
		os.Exit(1)
	}
	dbcomparator.CompareDB(*old, *new)
}
