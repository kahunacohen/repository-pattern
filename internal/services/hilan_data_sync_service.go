package services

import (
	"fmt"

	"github.com/kahunacohen/repo-pattern/internal/repositories"
)

type HilanDataSyncService struct {
	employeeRepo     repositories.EmployeeRepository
	familyStatusRepo repositories.FamilyStatusRepository
}

func NewHilanDataSyncService(employeeRepository repositories.EmployeeRepository, familyStatusRepository repositories.FamilyStatusRepository) *HilanDataSyncService {
	return &HilanDataSyncService{
		employeeRepo:     employeeRepository,
		familyStatusRepo: familyStatusRepository,
	}
}
func (ds *HilanDataSyncService) SyncRecords(records []hilanRecord) error {
	fmt.Println(len(records))
	return nil
}
