package services

import (
	"os"
	"strings"
	"testing"

	"github.com/kahunacohen/repo-pattern/db/generated"
)

func TestParseInputStreamToRecordsLineTooShort(t *testing.T) {
	reader := strings.NewReader("too\nshort line")
	hilanImporterService := HilanImportService{}
	_, err := hilanImporterService.ParseInputStreamToRecords(reader)
	if err.Error() != "error parsing line to record: error parsing line from hilan import file. length of line 1 is less than min length of 410" {
		t.Fatalf("did not get line too short error")
	}
}
func TestParseInputStreamToRecords(t *testing.T) {
	fh, err := os.Open("./MBTD594.PCF")
	if err != nil {
		t.Fatalf("faled to open test file: %v", err)
	}
	defer fh.Close()

	hilanImporterService := HilanImportService{familyStatuses: map[int]*generated.FamilyStatus{
		0: {ID: 1, Name: "single"},
		1: {ID: 2, Name: "married"},
		2: {ID: 3, Name: "devorce"},
		3: {ID: 4, Name: "widow"},
	}}
	records, err := hilanImporterService.ParseInputStreamToRecords(fh)
	if err != nil {
		t.Fatalf("failed to parse file: %v", err)
	}
	lenRecords := len(records)

	const inputFileLineCount = 2166
	if lenRecords != inputFileLineCount {
		t.Fatalf("wanted 1742, got: %d", lenRecords)
	}

	firstRecord := records[0]
	firstLocalID := firstRecord.LocalID
	if firstRecord.LocalID != "036003895" {
		t.Fatalf("wanted first localID to be '314125147', got: %s", firstLocalID)
	}
	// Check all local IDs are 9 numbers
	for i, r := range records {
		lenLocalID := len(r.LocalID)
		if lenLocalID != 9 {
			t.Fatalf("wanted length of 9 for localID '%s' at line %d, got: %d", r.LocalID, i+1, lenLocalID)
		}
	}
	if *firstRecord.City != "ניר עקיבא" {
		t.Fatalf("wanted 'ניר עקיבא', got %s", *firstRecord.City)
	}

	// jsonData, err := json.MarshalIndent(records, "", "  ") // Pretty print
	// if err != nil {
	// 	t.Fatalf("failed to marshal records to JSON: %v", err)
	// }
	// fmt.Println(string(jsonData)) // Print to standard output
}
