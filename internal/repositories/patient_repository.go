package repositories

import (
	"context"

	"github.com/kahunacohen/repo-pattern/db/generated"
)

type PatientRepository interface {
	GetOne(ctx context.Context, id int64) (*generated.Patient, error)
	LoadSQL(path string) error
}
