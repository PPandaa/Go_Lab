package pkg

import (
	"encoding/base64"
	"strings"
	"time"
)

// Encode
func Base64Encode(targetString string) string {

	encodeString := base64.StdEncoding.EncodeToString([]byte(targetString))
	return encodeString

}

// Time
func ConvertStringToTime(timeString string) time.Time {

	timeFormat, _ := time.Parse(time.RFC3339, timeString)
	return timeFormat

}

// String
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

// Type
func CheckType(i interface{}) string {

	switch i.(type) {
	case int:
		return "int"
	case float32:
		return "float32"
	case float64:
		return "float64"
	case string:
		return "string"
	default:
		return "none"
	}

}
