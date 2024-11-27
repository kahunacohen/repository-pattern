package services

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

var (
	// encoding = charmap.CodePage862.NewEncoder()
	decoder = charmap.CodePage862.NewDecoder()
)

type hilanRecord struct {
	// Birthday         *time.Time `json:"birthday"`
	City *string `json:"city"`
	// Email            string     `json:"email"`
	// EndDate          *time.Time `json:"endDate"`
	// FamilyStatus     *int64     `json:"familyStatus"`
	FirstName string `json:"firstName"`
	LocalID   string `json:"localID"`
	// Passport         string     `json:"password"`
	// PhoneNumber      *string    `json:"phoneNumber"`
	// PhoneNumber2     *string    `json:"phoneNumber2"`
	// SpouceFirstName  *string    `json:"spouceFirstName"`
	// StartWorkingDate *time.Time `json:"startWorkingDate"`
	// Status           *string    `json:"status"`
	Street  *string `json:"street"`
	Surname string  `json:"surname"`
	// Tarrif           string     `json:"tarrif"`
}

func parseInputStreamToRecords(r io.Reader) ([]hilanRecord, error) {
	var records []hilanRecord
	bufReader := bufio.NewReader(r)
	for {
		line, err := bufReader.ReadBytes('\n')
		if err != nil && err.Error() != "EOF" {
			return nil, fmt.Errorf("error reading bytes: %w", err)
		}
		if err == io.EOF {
			break
		}

		// Skip blank or empty lines
		if strings.TrimSpace(string(line)) == "" {
			continue
		}

		record, err := parseLineToRecord(line)
		if err != nil {
			return nil, fmt.Errorf("error parsing line to record: %w", err)
		}
		records = append(records, *record)
	}
	return records, nil
}

func parseLineToRecord(line []byte) (*hilanRecord, error) {
	var record hilanRecord
	buf := bytes.NewBuffer(line)
	// Skip first two 0s. "factory number"
	buf.Next(2)

	// Skip salary number
	buf.Next(6)

	// Read the next 9 to a string.
	localID := *readString(buf.Next(8)) + *readString(buf.Next(1))
	record.LocalID = localID

	surname := *readReverseString(buf.Next(15))
	firstName := *readReverseString(buf.Next(15))
	record.Surname = surname
	record.FirstName = firstName

	// salary department
	buf.Next(2)

	// address
	street := readReverseString(buf.Next(20))
	city := readReverseString(buf.Next(15))
	record.Street = street
	record.City = city

	return &record, nil
}

func readString(buffer []byte) *string {
	decoded, _, _ := transform.Bytes(decoder, buffer)
	str := strings.TrimSpace(string(decoded))
	return &str
}

func readReverseString(buffer []byte) *string {
	str := readString(buffer)
	r := []rune(*str)
	var res []rune
	for i := len(r) - 1; i >= 0; i-- {
		res = append(res, r[i])
	}
	resStr := string(res)
	return &resStr
}
