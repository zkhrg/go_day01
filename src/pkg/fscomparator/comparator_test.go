package fscomparator_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/zkhrg/go_day01/pkg/fscomparator"
)

func TestCompareFS(t *testing.T) {
	oldFileContent := "/etc/file1.txt\n" + "/etc/file3.txt"
	newFileContent := "/etc/file2.txt\n"
	expectedOutput := "REMOVED\t| /etc/file1.txt\n" +
		"REMOVED\t| /etc/file3.txt\n" +
		"ADDED\t| /etc/file2.txt\n"

	oldFile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(oldFile.Name())

	_, err = oldFile.WriteString(oldFileContent)
	if err != nil {
		t.Fatal(err)
	}

	if err := oldFile.Close(); err != nil {
		t.Fatal(err)
	}

	newFile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(newFile.Name())

	_, err = newFile.WriteString(newFileContent)
	if err != nil {
		t.Fatal(err)
	}

	if err := newFile.Close(); err != nil {
		t.Fatal(err)
	}

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}

	oldStdout := os.Stdout
	os.Stdout = w

	defer func() {
		os.Stdout = oldStdout
		w.Close()
	}()

	tokens_map := make(map[string]uint32)
	decode_tokens := make([]string, 0)
	old_fs := fscomparator.NewFileSystem(oldFile.Name(), &tokens_map, &decode_tokens)
	new_fs := fscomparator.NewFileSystem(newFile.Name(), &tokens_map, &decode_tokens)
	fscomparator.CompareFS(&old_fs, &new_fs)

	w.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("Failed to read from pipe: %v", err)
	}

	actualOutput := buf.String()
	if actualOutput != expectedOutput {
		t.Errorf("Expected output: %q, but got: %q", expectedOutput, actualOutput)
	}
}
