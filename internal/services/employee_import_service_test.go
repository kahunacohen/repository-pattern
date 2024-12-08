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

	hilanImportService := HilanImportService{familyStatuses: map[int]*generated.FamilyStatus{
		0: {ID: 1, Name: "single"},
		1: {ID: 2, Name: "married"},
		2: {ID: 3, Name: "devorce"},
		3: {ID: 4, Name: "widow"},
	}}
	records, err := hilanImportService.ParseInputStreamToRecords(fh)
	if err != nil {
		t.Fatalf("failed to parse file: %v", err)
	}
	lenRecords := len(records)

	// The test file has some lines with no working period, these return nil
	// records (3).
	const numNilRecords = 3
	const inputFileLineCount = 2166 - numNilRecords
	if lenRecords != inputFileLineCount {
		t.Fatalf("wanted %d, got: %d", inputFileLineCount, lenRecords)
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
	if *firstRecord.FamilyStatusId != 2 {
		t.Fatalf("wanted 2 for family status ID, got %d", *firstRecord.FamilyStatusId)
	}
	formattedBirthday := firstRecord.Birthday.Format("2006-01-02 15:04:05 +200 IST")

	if formattedBirthday != "1980-01-04 00:00:00 +400 IST" {
		t.Fatalf("wanted 1980-01-04 00:00:00 +0200 IST, got %s", formattedBirthday)

	}
	if *firstRecord.Mobile != "+972523438921" {
		t.Fatalf("wanted +972523438921, got %s", *firstRecord.Mobile)
	}
	if firstRecord.PhoneNumber2 != nil {
		t.Fatal("phone number 2 should be nil")
	}

	if firstRecord.Tarrif != "4261" {
		t.Fatalf("wanted 4261, got %s", firstRecord.Tarrif)
	}
	if *firstRecord.Status != "45" {
		t.Fatalf("wanted 45, got %s", *firstRecord.Status)
	}

	// The tenth record has a spouce first name
	if *records[10].SpouceFirstName != "רן" {
		t.Fatalf("wanted got '%s'", *records[10].SpouceFirstName)
	}

	if *firstRecord.Email != "1maayanf@matav.org.il" {
		t.Fatalf("wanted 1maayanf@matav.org.il, got %s", *firstRecord.Email)
	}

	// jsonData, err := json.MarshalIndent(records, "", "  ") // Pretty print
	// if err != nil {
	// 	t.Fatalf("failed to marshal records to JSON: %v", err)
	// }
	// fmt.Println(string(jsonData)) // Print to standard output
}
