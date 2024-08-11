package flaghelper_test

import (
	"errors"
	"testing"

	"github.com/zkhrg/go_day01/pkg/flaghelper"
)

func TestCheckFlagLenght(t *testing.T) {
	tests := []struct {
		osargs       []string
		checkargs    []string
		expected_err error
	}{
		{[]string{"--new", "file", "--old", "asd", "asd3"}, []string{"--new", "--old"}, nil},
		{[]string{"xcx", "xxx", "--f", "asd", "asd3"}, []string{"-f"}, errors.New("")},
		{[]string{"xcx", "xxx", "-old", "--new", "asd3"}, []string{"--new", "--old"}, errors.New("")},
		{[]string{"xcx", "xxx", "-old", "--new", "-f"}, []string{"-f"}, nil},
	}
	for _, test := range tests {
		if err := flaghelper.CheckLengthFlags(test.osargs, test.checkargs); err == nil && test.expected_err != nil || err != nil && test.expected_err == nil {
			t.Errorf("CheckFlagLenght: expected %s - have %s", err.Error(), test.expected_err.Error())
		}
	}
}
