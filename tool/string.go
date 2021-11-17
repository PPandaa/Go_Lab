package tool

import (
	"encoding/json"
	"strings"
)

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

func ConvertStringToMap(str string) map[string]interface{} {

	strMap := make(map[string]interface{})
	json.Unmarshal([]byte(str), &strMap)
	return strMap

}
