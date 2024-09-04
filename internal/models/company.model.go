package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	ID                 string `json:"id" gorm:"type:varchar(36);primaryKey"`
	CompanyName        string `json:"company_name" gorm:"type:varchar(255);not null;unique"`
	CompanyDescription string `json:"company_description" gorm:"type:varchar(255);not null"`
	CompanyLogo        string `json:"company_logo" gorm:"type:varchar(255);not null"`
}

func (c *Company) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
