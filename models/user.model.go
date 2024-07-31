package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       *uint   `json:"id" gorm:"primaryKey"`
	Name     *string `json:"name" gorm:"type:varchar(255);not null"`
	Address  *string `json:"address" gorm:"type:varchar(255);not null"`
	Email    *string `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password *string `json:"password" gorm:"type:varchar(255);not null"`
}
