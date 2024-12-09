package repositories

import (
	"context"

	"github.com/kahunacohen/repo-pattern/db/generated"
)

type CompanyRepo interface {
	GetFirst(ctx context.Context) (*generated.Company, error)
}
