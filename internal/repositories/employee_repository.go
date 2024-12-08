package repositories

import "github.com/kahunacohen/repo-pattern/db/generated"

type EmployeeRepository interface {
	GetEmployeeByLocalIdOrPassport(localId, passportNumber *string) (*generated.Employee, error)
}
