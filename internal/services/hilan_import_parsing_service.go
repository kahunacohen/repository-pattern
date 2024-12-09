// A service to handle importing Hilan export files to the matav database.

package services

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/mail"
	"strconv"
	"strings"
	"time"

	"github.com/kahunacohen/repo-pattern/db/generated"
	"github.com/ttacon/libphonenumber"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

const (
	MIN_HILAN_LINE_LENGTH = 410
)

var (
	decoder = charmap.CodePage862.NewDecoder()
)

type hilanRecord struct {
	Birthday        *time.Time `json:"birthday"`
	City            *string    `json:"city"`
	Email           *string    `json:"email"`
	FamilyStatus    *int64     `json:"familyStatus"`
	FirstName       string     `json:"firstName"`
	LocalID         string     `json:"localID"`
	Mobile          *string    `json:"mobile"`
	PhoneNumber2    *string    `json:"phoneNumber2"`
	SpouceFirstName *string    `json:"spouceFirstName"`
	Status          *string    `json:"status"`
	Street          *string    `json:"street"`
	Surname         string     `json:"surname"`
	Tarrif          string     `json:"tarrif"`
}

type HilanImportParsingService struct {
	familyStatuses map[int]*generated.FamilyStatus
}

func (h *HilanImportParsingService) ParseStream(r io.Reader) ([]hilanRecord, error) {
	var records []hilanRecord
	bufReader := bufio.NewReader(r)
	i := 0
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
		record, err := h.parseLine(i, line)
		if err != nil {
			return nil, fmt.Errorf("error parsing line to record: %w", err)
		}
		if record != nil {
			records = append(records, *record)
		} else {
			log.Printf("record %d is nil, skipping", i+1)
		}
		i += 1
	}
	return records, nil
}

func (h *HilanImportParsingService) parseLine(lineNum int, line []byte) (*hilanRecord, error) {
	if len(line) < MIN_HILAN_LINE_LENGTH {
		return nil, fmt.Errorf("error parsing line from hilan import file. length of line %d is less than min length of %d", lineNum+1, MIN_HILAN_LINE_LENGTH)
	}
	var record hilanRecord
	buf := bytes.NewBuffer(line)
	// Skip first two 0s. "factory number"
	buf.Next(2)

	// Skip salary number
	buf.Next(6)

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

	// family status
	familyBuf := buf.Next(1)
	familyStatusFromFile, err := readInt64(familyBuf)
	if err != nil {
		// @TODO what about this error?
		log.Printf("error parsing line %d. Cannot read family status \"%v\" from file\n", lineNum, string(familyBuf))
	}
	record.FamilyStatus = familyStatusFromFile

	birthday, err := readDate(buf.Next(8), "20060102")
	if err != nil {
		return nil, fmt.Errorf("couldn't parse birthday")
	}
	record.Birthday = birthday

	//num kids
	buf.Next(2)

	// partner id number
	buf.Next(8)
	buf.Next(1)

	// partner works
	buf.Next(1)

	// bank number
	buf.Next(2)

	// bank branch number
	buf.Next(3)

	// account number
	buf.Next(9)

	startWorkingDate, err := readDate(buf.Next(8), "20060102")
	if err != nil {
		return nil, errors.New("start working period")
	}
	if startWorkingDate == nil {
		return nil, nil
	}

	// immigration date
	buf.Next(6)

	// derug name
	buf.Next(4)
	buf.Next(3)
	buf.Next(2)

	// phone number
	phoneNumber := readPhoneNumber(buf.Next(10))
	phoneNumber2 := readPhoneNumber(buf.Next(10))
	if phoneNumber != nil && *phoneNumber != "0000000000" {
		record.Mobile = phoneNumber
	}
	if phoneNumber2 != nil && *phoneNumber2 != "0000000000" {
		record.PhoneNumber2 = phoneNumber2
	}

	// religious
	buf.Next(1)

	// medical approval
	buf.Next(2)

	// caregiver course
	buf.Next(2)

	// has car
	buf.Next(2)

	// has driving license
	_ = *readString(buf.Next(2))

	// passport
	buf.Next(15)
	// several fields
	buf.Next(78)
	tariffBuf := buf.Next(10)
	tariffStr := strings.Trim(*readString(tariffBuf), " ")
	record.Tarrif = tariffStr

	// hospital
	buf.Next(4)

	// Status
	record.Status = readString(buf.Next(2))

	// percentage
	buf.Next(4)

	buf.Next(9)
	buf.Next(9)
	buf.Next(35)

	record.SpouceFirstName = readReverseString(buf.Next(10))

	buf.Next(14)

	setEmail(buf, &record)

	buf.Next(14)

	return &record, nil
}

// This function takes a family status number from the hilan file, converts it to
// Matav's family status db and sets the record with that ID.
//
//	func (h *HilanImportParsingService) setFamilyStatusID(record *hilanRecord, familyStatusFromFile *int64) {
//		if familyStatusFromFile != nil {
//			familyStatusValueFromFile := *familyStatusFromFile
//			if familyStatusValueFromFile > 5 {
//				// @TODO why minus 5 from values over 5?
//				familyStatusValueFromFile = familyStatusValueFromFile - 5
//			}
//			// the family statuses map is a map of accounting IDs to FamilyStatus structs, each
//			// with the Database ID and the name (e.g. single, married etc). I assume the status in the file
//			// is the accounting ID?
//			// DB looks like this:
//			// id |  name   | accounting_id
//			// ----+---------+---------------
//			//   1 | single  |             0
//			//   2 | married |             0
//			//   3 | devorce |             0
//			//   4 | widow   |             0
//			// @TODO the odd thing is that in our real db, every family status has an accounting_id of 0.
//			// For our test we inject a map, but we can't statically assign a map with identical keys of 0.
//			// This means that in the real job, it will always be single...?!!
//			familyStatus, ok := h.familyStatuses[int(familyStatusValueFromFile)]
//			if ok {
//				record.FamilyStatusId = &familyStatus.ID
//			}
//		}
//	}
func setEmail(buf *bytes.Buffer, record *hilanRecord) {
	email := strings.Trim(string(buf.Next(31)), " ")
	if len(email) == 0 {
		record.Email = nil
	} else {
		addr, err := mail.ParseAddress(email)
		if err != nil {
			record.Email = nil
		} else {
			record.Email = &addr.Address
		}
	}
}

func readPhoneNumber(buffer []byte) *string {
	phoneStr := readString(buffer)
	if phoneStr == nil {
		return nil
	}
	if *phoneStr == "0000000000" {
		return nil
	}

	num, err := libphonenumber.Parse(fmt.Sprintf("972%v", *phoneStr), "IL")
	if err != nil {
		return phoneStr
	}
	formatted := libphonenumber.Format(num, libphonenumber.E164)
	return &formatted
}
func readDate(buffer []byte, format string) (*time.Time, error) {
	str := *readString(buffer)
	if str == "" {
		return nil, nil
	}
	timezone, err := time.LoadLocation("Asia/Jerusalem")
	if err != nil {
		return nil, err
	}
	d, err := time.ParseInLocation(format, str, timezone)
	return &d, err
}

func readInt64(buffer []byte) (*int64, error) {
	num, err := strconv.ParseInt(*readString(buffer), 10, 64)
	return &num, err
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
