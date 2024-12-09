package repositories

import (
	"database/sql"

	"context"

	"github.com/kahunacohen/repo-pattern/db/generated"
)

type FamilyStatusImpl struct {
	DB *sql.DB
}

func (f *FamilyStatusImpl) GetAll(ctx context.Context) ([]generated.FamilyStatus, error) {
	return nil, nil
}
