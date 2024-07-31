package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	ID                 *uint `json:"id" gorm:"primaryKey"`
	CompanyID          *int
	Company            *Company `json:"company" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RoleID             *int
	Role               *Role   `json:"role" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProjectTitle       *string `json:"project_title" gorm:"type:varchar(255);not null"`
	ProjectDescription *string `json:"project_description" gorm:"type:text;not null"`
	JobType            *string `json:"job_type" gorm:"type:enum('fulltime','parttime','freelance');not null"`
	Status             *string `json:"status" gorm:"type:enum('open','closed');not null"`
}
