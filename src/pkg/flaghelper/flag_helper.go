package flaghelper

import "errors"

func CheckLengthFlags(osargs []string, checkargs []string) error {
	found := true
	for i := 0; i < len(checkargs) && found; i++ {
		mini_found := false
		for j := 0; j < len(osargs); j++ {
			if checkargs[i] == osargs[j] {
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
