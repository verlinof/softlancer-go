package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Application struct {
	ID             string `json:"id" gorm:"type:varchar(36);primaryKey"`
	UserID         string
	User           User `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProjectID      string
	Project        Project `json:"project" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CuriculumVitae string  `json:"curiculum_vitae" gorm:"type:varchar(255);not null"`
	Portofolio     string  `json:"portofolio" gorm:"type:varchar(255);not null"`
	Status         string  `json:"status" gorm:"default:'waiting';type:enum('waiting','accepted','rejected');not null"`
}

func (application *Application) BeforeCreate(tx *gorm.DB) (err error) {
	application.ID = uuid.New().String() // Generate a new UUID
	return
}
