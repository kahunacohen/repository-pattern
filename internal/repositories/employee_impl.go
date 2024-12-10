package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/kahunacohen/repo-pattern/db/generated"
)

type EmployeeRepositoryImpl struct {
	DB *sql.DB
}

// Returns nil, nil if the record is not found, decoupling this method from the underlying db.
func (e *EmployeeRepositoryImpl) GetEmployeeByLocalIdOrPassport(ctx context.Context, localId, passportNumber *string) (*generated.Employee, error) {
	queries := generated.New(e.DB)
	localIdNullStr := ToSqlNullStr(localId)
	if localIdNullStr.Valid {
		localIdNullStr.String = strings.TrimLeft(localIdNullStr.String, "0")
	}
	empl, err := queries.GetEmployeeByLocalIdOrPassport(ctx, generated.GetEmployeeByLocalIdOrPassportParams{
		LocalIDNumber:         ToSqlNullStr(localId),
		ForeignPassportNumber: ToSqlNullStr(passportNumber),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting employee by local ID or passport: %w", err)
	}
	return &empl, nil
}

func (e *EmployeeRepositoryImpl) UpdateEmployee(ctx context.Context, params generated.UpdateEmployeeParams) error {
	queries := generated.New(e.DB)
	err := queries.UpdateEmployee(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func ToSqlNullStr(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	} else {
		return sql.NullString{Valid: true, String: *s}
	}
}
