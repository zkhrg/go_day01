package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Path struct {
	tag  string
	path map[string]Path
}

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
	compareFS(*old, *new)
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

func compareFS(old_db_filename, new_db_filename string) {
	var old_tree, new_tree Path
	old_tree.FillFromFile(old_db_filename)
	new_tree.FillFromFile(new_db_filename)
	Print(old_tree)
}

// type Path struct {
// 	path map[string]Path
// }

func (s *Path) Add(filename string) {
	root := *s
	parts := strings.Split(filename, "/")
	for _, part := range parts {
		if part == "" {
			continue
		}
		fpart := fmt.Sprintf("/%s", part)
		if _, ok := root.path[fpart]; !ok {
			root.path[fpart] = NewPath(fpart)
		}
		root = root.path[fmt.Sprintf("/%s", part)]
	}
}

func (s *Path) FillFromFile(filename string) {
	*s = NewPath("")
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text() // Получаем текущую строку
		s.Add(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}
}

func NewPath(fpart string) Path {
	return Path{
		path: make(map[string]Path),
		tag:  fpart,
	}
}

// func (s *Path) Print() {
// 	for
// }

func Print(p Path) {
	if len(p.path) == 0 {
		fmt.Printf("%s\n", p.tag)
	}
	for _, v := range p.path {
		Print(v)
	}
}
