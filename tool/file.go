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

func WriteJsonFile(file_name string, data interface{}) {
	// Open a file for writing
	json_file, err := os.Create(file_name)
	if err != nil {
		guard.Logger.Fatal("Error creating file - " + err.Error())
		return
	}
	defer json_file.Close()

	// Create a JSON encoder and write the struct to the file
	json_file_encoder := json.NewEncoder(json_file)
	json_file_encoder.SetIndent("", "  ")  // Pretty print with indentation
	json_file_encoder.SetEscapeHTML(false) // Disable escaping of &, <, >
	if err := json_file_encoder.Encode(data); err != nil {
		guard.Logger.Fatal("Error encoding JSON - " + err.Error())
	}
}
