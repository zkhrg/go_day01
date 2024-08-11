package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/zkhrg/go_day01/pkg/fscomparator"
)

func main() {
	old := flag.String("old", "", "old version db")
	new := flag.String("new", "", "new version db")

	flag.Parse()

	checkLengthFlags(os.Args, []string{"--new", "--old"})

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
	old_fs := fscomparator.NewFileSystem(*old)
	new_fs := fscomparator.NewFileSystem(*new)

	old_fs.Tokens_map = &tokens_map
	new_fs.Tokens_map = &tokens_map
	old_fs.Decode_tokens = &decode_tokens
	new_fs.Decode_tokens = &decode_tokens

	old_fs.Fill()
	new_fs.Fill()

	fscomparator.CompareFS(&old_fs, &new_fs)
}

func checkLengthFlags(osargs []string, checkargs []string) error {
	found := true
	for i := 0; i < len(checkargs) && found; i++ {
		mini_found := false
		for j := 0; j < len(osargs); j++ {
			if checkargs[i] == osargs[i] {
				mini_found = true
			}
		}
		found = found && mini_found
	}
	if found {
		return nil
	}
	return errors.New("not found right version of flags")
}
