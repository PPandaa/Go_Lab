package pkg

import "strings"

func IsEmptyString(targetString string) bool {

	if len(strings.TrimSpace(targetString)) == 0 {
		return true
	} else {
		return false
	}

}
