package usecase

import (
	"context"
	"patientdx-backend-web-go/domain"
	"time"
)

type patientUseCase struct {
	patientRepository domain.PatientRepository
	contextTimeout    time.Duration
}

func NewPatientUseCase(patient domain.PatientRepository, t time.Duration) domain.PatientUseCase {
	return &patientUseCase{
		patientRepository: patient,
		contextTimeout:    t,
	}
}

func (c *patientUseCase) FetchPatientByID(ctx context.Context, id uint) (*domain.Patient, error) {
	res, err := c.patientRepository.RetrievePatientByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *patientUseCase) FetchPatients(ctx context.Context) ([]domain.Patient, error) {
	res, err := c.patientRepository.RetrieveAllPatient()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *patientUseCase) CreatePatient(ctx context.Context, req *domain.Patient) (*domain.Patient, error) {
	res, err := c.patientRepository.CreatePatient(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *patientUseCase) UpdatePatient(ctx context.Context, req *domain.Patient) (*domain.Patient, error) {
	res, err := c.patientRepository.UpdatePatient(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *patientUseCase) DeletePatient(ctx context.Context, id uint) error {
	err := c.patientRepository.DeletePatient(id)
	if err != nil {
		return err
	}
	return nil
}
