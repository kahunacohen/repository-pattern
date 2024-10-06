package services

import (
	"context"
	"testing"

	"github.com/kahunacohen/repo-pattern/internal/repositories"
)

func TestGetPatientFromSqlite(t *testing.T) {
	repo, err := repositories.NewSqliteRepository("../../demo.db")
	if err != nil {
		t.Fail()
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
