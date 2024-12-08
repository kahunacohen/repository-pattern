package repositories

import (
	"context"

	"github.com/kahunacohen/repo-pattern/db/generated"
)

type EmployeeRepository interface {
	GetEmployeeByLocalIdOrPassport(ctx context.Context, localId, passportNumber *string) (*generated.Employee, error)
}
