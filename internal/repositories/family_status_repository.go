package repositories

import (
	"context"

	"github.com/kahunacohen/repo-pattern/db/generated"
)

type FamilyStatusRepository interface {
	GetAll(ctx context.Context) ([]generated.FamilyStatus, error)
}
