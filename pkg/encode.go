package pkg

import "encoding/base64"

func Base64Encode(targetString string) string {
	encodeString := base64.StdEncoding.EncodeToString([]byte(targetString))
	return encodeString
}
