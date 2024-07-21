package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Id *int `json:"id" gorm:"primaryKey"`
	CompanyName *string `json:"company_name" gorm:"type:varchar(255);not null"`
	CompanyDescription *string `json:"company_description" gorm:"type:varchar(255);not null"`
	CompanyLogo *string `json:"company_logo" gorm:"type:varchar(255);not null"`
}