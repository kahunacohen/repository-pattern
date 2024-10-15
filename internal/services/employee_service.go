package services

import "time"

type hilanImportFile struct {
	Birthday         *time.Time `json:"birthday"`
	City             *string    `json:"city"`
	Email            string     `json:"email"`
	EndDate          *time.Time `json:"endDate"`
	FamilyStatus     *int64     `json:"familyStatus"`
	FirstName        string     `json:"firstName"`
	LocalID          string     `json:"localID"`
	Passport         string     `json:"password"`
	PhoneNumber      *string    `json:"phoneNumber"`
	PhoneNumber2     *string    `json:"phoneNumber2"`
	SpouceFirstName  *string    `json:"spouceFirstName"`
	StartWorkingDate *time.Time `json:"startWorkingDate"`
	Status           *string    `json:"status"`
	Street           *string    `json:"street"`
	Surname          string     `json:"surname"`
	Tarrif           string     `json:"tarrif"`
}
