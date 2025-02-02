package tool

import (
	"GoLab/guard"
	"encoding/json"
	"io"
	"os"
)

func ReadJsonFile(file_name string) map[string]interface{} {
	// Open json_file
	json_file, err := os.Open(file_name)
	if err != nil {
		guard.Logger.Fatal("Loading JSON file Error - " + err.Error())
	}

	byte_value, _ := io.ReadAll(json_file)

	var json_file_content map[string]interface{}
	json.Unmarshal([]byte(byte_value), &json_file_content)

	// Closing json_file
	json_file.Close()

	return json_file_content
}
