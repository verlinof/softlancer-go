package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Id                *int    `json:"id" gorm:"primaryKey"`
	CompanyId         *int    `json:"company_id" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Company           *Company `json:"company" gorm:"foreignKey:CompanyId"`
	RoleId            *int    `json:"role_id" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Role              *Role   `json:"role" gorm:"foreignKey:RoleId"`
	ProjectTitle      *string `json:"project_title" gorm:"type:varchar(255);not null"`
	ProjectDescription *string `json:"project_description" gorm:"type:text;not null"`
	JobType           *string `json:"job_type" gorm:"type:enum('fulltime','parttime','freelance');not null"`
	Status            *string `json:"status" gorm:"type:enum('open','closed');not null"`
}
