package services

import (
	"fmt"

	"github.com/kahunacohen/repo-pattern/internal/repositories"
)

type HilanDataSyncService struct {
	employeeRepo     repositories.EmployeeRepo
	familyStatusRepo repositories.FamilyStatusRepo
}

func NewHilanDataSyncService(employeeRepo repositories.EmployeeRepo, familyStatusRepo repositories.FamilyStatusRepo) *HilanDataSyncService {
	return &HilanDataSyncService{
		employeeRepo:     employeeRepo,
		familyStatusRepo: familyStatusRepo,
	}
}
func (ds *HilanDataSyncService) SyncRecords(records []hilanRecord) error {
	fmt.Println(len(records))
	return nil
}
