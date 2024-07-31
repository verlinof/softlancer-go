package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID       *int    `json:"id" gorm:"primaryKey"`
	RoleName *string `json:"role_name" gorm:"type:varchar(255);not null"`
}
