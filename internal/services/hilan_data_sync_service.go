package services

import (
	"fmt"

	"github.com/kahunacohen/repo-pattern/internal/repositories"
)

type HilanDataSyncService struct {
	employeeRepo repositories.EmployeeRepository
}

func NewHilanDataSyncService(employeeRepository repositories.EmployeeRepository) *HilanDataSyncService {
	return &HilanDataSyncService{
		employeeRepo: employeeRepository,
	}
}
func (ds *HilanDataSyncService) SyncRecords(records []hilanRecord) error {
	fmt.Println(len(records))
	return nil
}
