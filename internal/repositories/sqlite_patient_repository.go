package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/kahunacohen/repo-pattern/db/generated"
	_ "github.com/mattn/go-sqlite3"
)

type SqlitePatientRepository struct {
	db *sql.DB
}

func NewSqliteRepository(dbPath string) (*SqlitePatientRepository, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return &SqlitePatientRepository{db: db}, nil
}

func (repo *SqlitePatientRepository) GetOne(ctx context.Context, id int64) (*generated.Patient, error) {
	queries := generated.New(repo.db)
	p, err := queries.GetPatient(ctx, id)
	return &p, err
}

func (repo *SqlitePatientRepository) LoadSQL(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("repository error loading schema: %w", err)
	}
	_, err = repo.db.Exec(string(data))
	return err
}
