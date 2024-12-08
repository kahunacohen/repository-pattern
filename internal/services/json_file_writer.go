package services

import (
	"encoding/json"
	"os"
)

type JSONFileWriter struct{}

func (j *JSONFileWriter) WriteToFile(filename string, data interface{}) error {
	// Open the file for writing (create if not exists, truncate if it does).
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a JSON encoder for the file.
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty-printing (optional)

	// Write the data to the file.
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}
