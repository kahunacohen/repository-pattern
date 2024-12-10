package services

import (
	"context"
	"fmt"

	"github.com/kahunacohen/repo-pattern/db/generated"
	"github.com/kahunacohen/repo-pattern/internal/repositories"
)

type HilanDataSyncService struct {
	ctx              context.Context
	employeeRepo     repositories.EmployeeRepo
	familyStatusRepo repositories.FamilyStatusRepo
	familyStatuses   map[int]*generated.FamilyStatus
}

func NewHilanDataSyncService(ctx context.Context, companyRepo repositories.CompanyRepo, employeeRepo repositories.EmployeeRepo, familyStatusRepo repositories.FamilyStatusRepo) (*HilanDataSyncService, error) {
	company, err := companyRepo.GetFirst(ctx)
	if !company.EmployeeSyncActive {
		return nil, fmt.Errorf("error intializing HilanDataSyncService: company \"%s\" employee_sync_active flag not set", company.Name)
	}
	if err != nil {
		return nil, fmt.Errorf("error initializing HilanDataSyncService: %w", err)
	}
	familyStatuses, err := familyStatusRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing HilanDataSyncService: %w", err)
	}

	familyStatusesMap := make(map[int]*generated.FamilyStatus, len(familyStatuses))
	for _, fs := range familyStatuses {
		fstatus := fs
		familyStatusesMap[int(fs.AccountingID)] = &fstatus
	}

	return &HilanDataSyncService{
		ctx:              ctx,
		employeeRepo:     employeeRepo,
		familyStatusRepo: familyStatusRepo,
		familyStatuses:   familyStatusesMap,
	}, nil
}
func (ds *HilanDataSyncService) SyncRecords(records []hilanRecord) error {
	// for _, record := range records {
	// 	// ds.employeeRepo.GetEmployeeByLocalIdOrPassport(ds.ctx, &record.LocalID, "")
	// 	// fmt.Println(record)
	// }
	return nil
}
