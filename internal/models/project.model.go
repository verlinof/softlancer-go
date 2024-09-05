package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	ID                 string `json:"id" gorm:"type:varchar(36);primaryKey"`
	CompanyID          string
	Company            Company `json:"company" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RoleID             string
	Role               Role   `json:"role" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProjectTitle       string `json:"project_title" gorm:"type:varchar(255);not null"`
	ProjectDescription string `json:"project_description" gorm:"type:text;not null"`
	JobType            string `json:"job_type" gorm:"type:enum('fulltime','parttime','freelance');not null"`
	Status             string `json:"status" gorm:"type:enum('open','closed');not null"`
}

func (p *Project) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	return
}
