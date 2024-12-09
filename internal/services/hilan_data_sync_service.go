package services

import (
	"context"
	"fmt"

	"github.com/kahunacohen/repo-pattern/db/generated"
	"github.com/kahunacohen/repo-pattern/internal/repositories"
)

type HilanDataSyncService struct {
	companyRepo      repositories.CompanyRepo
	ctx              context.Context
	employeeRepo     repositories.EmployeeRepo
	familyStatusRepo repositories.FamilyStatusRepo
	familyStatuses   map[int]*generated.FamilyStatus
}

func NewHilanDataSyncService(ctx context.Context, companyRepo repositories.CompanyRepo, employeeRepo repositories.EmployeeRepo, familyStatusRepo repositories.FamilyStatusRepo) (*HilanDataSyncService, error) {
	company, err := companyRepo.GetFirst(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing HilanDataSyncService: %w", err)
	}
	fmt.Println(company)

	return &HilanDataSyncService{
		ctx:              ctx,
		employeeRepo:     employeeRepo,
		familyStatusRepo: familyStatusRepo,
	}, nil
}
func (ds *HilanDataSyncService) SyncRecords(records []hilanRecord) error {
	fmt.Println(len(records))
	return nil
}
