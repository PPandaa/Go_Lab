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
