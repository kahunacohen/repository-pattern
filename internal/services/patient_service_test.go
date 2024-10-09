package services

import (
	"context"
	"testing"

	"github.com/kahunacohen/repo-pattern/internal/repositories"
)

func TestGetPatientFromSqlite(t *testing.T) {
	repo, err := repositories.NewSqliteRepository(":memory:")
	if err != nil {
		t.Fatalf("error initializing repository: %v", err)
	}
	if err := LoadSchemaAndSeed(repo); err != nil {
		t.Fatalf("error loading schema and seeding: %v", err)
	}
	service := NewPatientService(repo)
	patient, err := service.GetPatient(context.Background(), 1)
	if err != nil {
		t.Fatalf("error getting patient: %v", err)
	}
	if patient.LocalID != "341077656" {
		t.Fatalf("wanted '341077656', got '%s'", patient.LocalID)
	}
}

func LoadSchemaAndSeed(repo *repositories.SqlitePatientRepository) error {
	if err := repo.LoadSQL("../../db/schema.sql"); err != nil {
		return err
	}
	if err := repo.LoadSQL("../../db/seed.sql"); err != nil {
		return err
	}
	return nil
}
