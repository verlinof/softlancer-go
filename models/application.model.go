package models

import "gorm.io/gorm"

type Application struct {
	gorm.Model
	ID             *uint `json:"id" gorm:"primaryKey"`
	UserID         *int
	User           *User `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProjectID      *int
	Project        *Project `json:"project" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CuriculumVitae *string  `json:"curiculum_vitae" gorm:"type:varchar(255);not null"`
	Portofolio     *string  `json:"portofolio" gorm:"type:varchar(255);not null"`
	Status         *string  `json:"status" gorm:"type:enum('waiting','accepted','rejected');not null"`
}
