package services

import (
	"context"
	"fmt"
	"log"

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
	if err != nil {
		return nil, fmt.Errorf("error initializing HilanDataSyncService: %w", err)
	}
	if !company.EmployeeSyncActive {
		return nil, fmt.Errorf("error intializing HilanDataSyncService: company \"%s\" employee_sync_active flag not set", company.Name)
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
func (h *HilanDataSyncService) SyncRecords(records []hilanRecord) error {
	for _, record := range records {
		employee, err := h.employeeRepo.GetEmployeeByLocalIdOrPassport(h.ctx, &record.LocalID, record.Passport)
		if employee == nil && err == nil {
			// not found
			log.Printf("employee %s not found by local ID or passport", record.LocalID)
			continue
		}
		if err != nil {
			return fmt.Errorf("error getting employee by local ID: %w", err)
		}

		// At this point we have a matching employee from the database.
		// Initialize params we will send to UpdateEmployee.
		updateParams := generated.UpdateEmployeeParams{}
		h.setFamilyStatusParam(record, &updateParams)

		return h.employeeRepo.UpdateEmployee(h.ctx, updateParams)

	}
	return nil
}

func (h *HilanDataSyncService) setFamilyStatusParam(record hilanRecord, params *generated.UpdateEmployeeParams) {
	if record.FamilyStatus != nil {
		status := *record.FamilyStatus
		if status > 5 {
			status = status - 5
		}
		selected, ok := h.familyStatuses[int(status)]
		if ok {
			params.FamilyStatusID = selected.ID
		}
	}
}
