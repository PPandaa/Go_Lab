package tool

import "strings"

func IsEmptyString(targetString string) bool {

	if len(strings.TrimSpace(targetString)) == 0 {
		return true
	} else {
		return false
	}

}

func IsStringExist(target string, array []string) bool {

	for _, item := range array {
		if item == target {
			return true
		}
	}
	return false

}
