package repository

import (
	"assyarif-backend-web-go/domain"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type posgrePatientRepository struct {
	DB *gorm.DB
}

func NewPostgrePatient(client *gorm.DB) domain.PatientRepository {
	return &posgrePatientRepository{
		DB: client,
	}
}

func (a *posgrePatientRepository) RetrieveAllPatient() ([]domain.Patient, error) {
	var res []domain.Patient
	err := a.DB.
		Model(domain.Patient{}).
		Find(&res).Error
	if err != nil {
		return []domain.Patient{}, err
	}
	fmt.Println(res)
	return res, nil
}

func (a *posgrePatientRepository) RetrievePatientByID(id uint) (*domain.Patient, error) {
	var res domain.Patient
	err := a.DB.
		Model(domain.Patient{}).
		Where("id = ?", id).
		Take(&res).Error
	if err != nil {
		return &domain.Patient{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Patient{}, fmt.Errorf("record not found")
	}
	fmt.Println(res)
	return &res, nil
}

func (a *posgrePatientRepository) CreatePatient(patient *domain.Patient) (*domain.Patient, error) {
	err := a.DB.
		Model(domain.Patient{}).
		Create(patient).Error
	if err != nil {
		return &domain.Patient{}, err
	}
	fmt.Println(patient)
	return patient, nil
}

func (a *posgrePatientRepository) UpdatePatient(patient *domain.Patient) (*domain.Patient, error) {
	err := a.DB.
		Model(domain.Patient{}).
		Where("id = ?", patient.ID).
		Updates(patient).Error
	if err != nil {
		return &domain.Patient{}, err
	}
	fmt.Println(patient)
	return patient, nil
}

func (a *posgrePatientRepository) DeletePatient(id uint) error {
	err := a.DB.
		Model(domain.Patient{}).
		Where("id = ?", id).
		Delete(&domain.Patient{}).Error
	if err != nil {
		return err
	}
	return nil
}
