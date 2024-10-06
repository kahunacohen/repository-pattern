package services

import (
	"context"

	"github.com/kahunacohen/repo-pattern/db/generated"
	"github.com/kahunacohen/repo-pattern/internal/repositories"
)

type PatientService struct {
	repo repositories.PatientRepository
}

func NewPatientService(repo repositories.PatientRepository) *PatientService {
	return &PatientService{repo: repo}
}

func (s *PatientService) GetPatient(ctx context.Context, id int64) (*generated.Patient, error) {
	return s.repo.GetOne(ctx, id)
}
