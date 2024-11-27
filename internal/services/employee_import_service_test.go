package services

import (
	"os"
	"testing"
)

func TestParseInputStreamToRecords(t *testing.T) {
	fh, err := os.Open("./MBTD594.PCF")
	if err != nil {
		t.Fatalf("faled to open test file: %v", err)
	}
	defer fh.Close()
	records, err := parseInputStreamToRecords(fh)
	if err != nil {
		t.Fatalf("failed to parse file: %v", err)
	}
	lenRecords := len(records)

	const inputFileLineCount = 2166
	if lenRecords != inputFileLineCount {
		t.Fatalf("wanted 1742, got: %d", lenRecords)
	}
	firstLocalID := records[0].LocalID
	if firstLocalID != "036003895" {
		t.Fatalf("wanted first localID to be '314125147', got: %s", firstLocalID)
	}

	for i, r := range records {
		lenLocalID := len(r.LocalID)
		if lenLocalID != 9 {
			t.Fatalf("wanted length of 9 for localID '%s' at line %d, got: %d", r.LocalID, i+1, lenLocalID)
		}
	}

	if *records[0].City != "ניר עקיבא" {
		t.Fatalf("wanted 'ניר עקיבא', got %s", *records[0].City)
	}

	// Convert JSON byte slice to string
	// jsonString := string(jsonData)

	// Print the JSON string
	// fmt.Println(jsonString)

}
