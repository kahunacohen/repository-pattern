package repositories

import (
	"context"

	"github.com/kahunacohen/repo-pattern/db/generated"
)

type EmployeeRepo interface {
	GetEmployeeByLocalIdOrPassport(ctx context.Context, localId, passportNumber *string) (*generated.Employee, error)
	UpdateEmployee(ctx context.Context, params generated.UpdateEmployeeParams) error
}
