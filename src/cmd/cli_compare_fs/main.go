package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zkhrg/go_day01/pkg/flaghelper"
	"github.com/zkhrg/go_day01/pkg/fscomparator"
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

	tokens_map := make(map[string]uint32)
	decode_tokens := make([]string, 0)
	old_fs := fscomparator.NewFileSystem(*old, &tokens_map, &decode_tokens)
	new_fs := fscomparator.NewFileSystem(*new, &tokens_map, &decode_tokens)

	fscomparator.CompareFS(&old_fs, &new_fs)
}
