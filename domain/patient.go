package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	ID                  uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Name                string         `gorm:"not null" json:"name"`
	MedicalRecordNumber string         `gorm:"not null" json:"medical_record_number"`
	MedicalDiagnosis    string         `gorm:"not null" json:"medical_diagnosis"`
	NursingDiagnosis    string         `gorm:"not null" json:"nursing_diagnosis"`
	AttendingPhysician  string         `gorm:"not null" json:"attending_physician"`
	HealthHistory       string         `gorm:"not null" json:"health_history"`
	PhysicalExamination string         `gorm:"not null" json:"physical_examination"`
	TherapeuticAction   string         `gorm:"not null" json:"therapeutic_action"`
	CreatedAt           *time.Time     `json:"created_at"`
	UpdatedAt           *time.Time     `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type PatientRepository interface {
	RetrieveAllPatient() ([]Patient, error)
	RetrievePatientByID(id uint) (*Patient, error)
	CreatePatient(Patient *Patient) (*Patient, error)
	UpdatePatient(Patient *Patient) (*Patient, error)
	DeletePatient(id uint) error
}

type PatientUseCase interface {
	FetchPatients(ctx context.Context) ([]Patient, error)
	FetchPatientByID(ctx context.Context, id uint) (*Patient, error)
	CreatePatient(ctx context.Context, req *Patient) (*Patient, error)
	UpdatePatient(ctx context.Context, req *Patient) (*Patient, error)
	DeletePatient(ctx context.Context, id uint) error
}
