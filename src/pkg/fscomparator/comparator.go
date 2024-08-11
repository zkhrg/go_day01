package fscomparator

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FileSystem struct {
	Filename      string
	Paths         [][]uint32
	Tokens_map    *map[string]uint32
	Decode_tokens *[]string
}

func NewFileSystem(filename string, tokens_map *map[string]uint32, decode_tokens *[]string) FileSystem {
	fs := FileSystem{
		Filename:      filename,
		Paths:         make([][]uint32, 0),
		Tokens_map:    tokens_map,
		Decode_tokens: decode_tokens,
	}
	fs.fill()
	return fs
}

func CompareFS(old_fs *FileSystem, new_fs *FileSystem) {
	for i, path := range old_fs.Paths {
		res := findSubslice(new_fs.Paths, path)
		if res != -1 {
			new_fs.Paths[res] = []uint32{}
			old_fs.Paths[i] = []uint32{}
		}
	}
	printDiff(old_fs, new_fs)
}

func (s *FileSystem) addPath(path string) {
	splitted_path := strings.Split(path, "/")
	compressed_path := make([]uint32, 0)
	for _, path_part := range splitted_path {
		if path_part == "" {
			continue
		}
		if _, ok := (*s.Tokens_map)[path_part]; !ok {
			(*s.Tokens_map)[path_part] = uint32(len(*s.Decode_tokens))
			*s.Decode_tokens = append(*s.Decode_tokens, path_part)
		}
		compressed_path = append(compressed_path, (*s.Tokens_map)[path_part])
	}
	s.Paths = append(s.Paths, compressed_path)
}

func (s *FileSystem) fill() {
	file, err := os.Open(s.Filename)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		s.addPath(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}
}

func slicesEqual(a, b []uint32) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func findSubslice(sliceOfSlices [][]uint32, target []uint32) int {
	for i, subSlice := range sliceOfSlices {
		if slicesEqual(subSlice, target) {
			return i
		}
	}
	return -1
}

func printDiff(old_fs *FileSystem, new_fs *FileSystem) {
	for _, v := range old_fs.Paths {
		if len(v) == 0 {
			continue
		}
		fpath := ""
		for i := 0; i < len(v); i++ {
			fpath += "/" + (*old_fs.Decode_tokens)[int(v[i])]
		}
		fmt.Printf("REMOVED\t| %s\n", fpath)
	}
	for _, v := range new_fs.Paths {
		if len(v) == 0 {
			continue
		}
		fpath := ""
		for i := 0; i < len(v); i++ {
			fpath += "/" + (*new_fs.Decode_tokens)[int(v[i])]
		}
		fmt.Printf("ADDED\t| %s\n", fpath)
	}
}
