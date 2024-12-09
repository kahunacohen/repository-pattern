package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kahunacohen/repo-pattern/db/generated"
)

type CompanyImpl struct {
	DB *sql.DB
}

func (c *CompanyImpl) GetFirst(ctx context.Context) (*generated.Company, error) {
	queries := generated.New(c.DB)
	company, err := queries.GetFirst(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting first company: %v", err)
	}
	return &company, nil
}
