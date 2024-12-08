package repositories

import (
	"database/sql"

	"github.com/kahunacohen/repo-pattern/db/generated"
)

type EmployeeRepositoryImpl struct {
	db *sql.DB
}

func (e *EmployeeRepositoryImpl) GetEmployeeByLocalIdOrPassport(localId, passportNumber *string) (*generated.Employee, error) {
	return nil, nil
}
