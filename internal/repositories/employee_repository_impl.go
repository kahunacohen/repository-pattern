package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/kahunacohen/repo-pattern/db/generated"
)

type EmployeeRepositoryImpl struct {
	db *sql.DB
}

func (e *EmployeeRepositoryImpl) GetEmployeeByLocalIdOrPassport(ctx context.Context, localId, passportNumber *string) (*generated.Employee, error) {
	queries := generated.New(e.db)
	localIdNullStr := ToSqlNullStr(localId)
	if localIdNullStr.Valid {
		localIdNullStr.String = strings.TrimLeft(localIdNullStr.String, "0")
	}
	empl, err := queries.GetByLocalIdOrPassport(ctx, generated.GetByLocalIdOrPassportParams{
		LocalIDNumber:         ToSqlNullStr(localId),
		ForeignPassportNumber: ToSqlNullStr(passportNumber),
	})
	if err != nil {
		return nil, fmt.Errorf("error getting employee by local ID or passport")
	}
	return &empl, nil
}

func ToSqlNullStr(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	} else {
		return sql.NullString{Valid: true, String: *s}
	}
}
