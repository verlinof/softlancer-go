package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       string `json:"id" gorm:"type:varchar(36);primaryKey"`
	Name     string `json:"name" gorm:"type:varchar(255);not null"`
	Address  string `json:"address" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
	IsAdmin  bool   `json:"is_admin" gorm:"default:false"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New().String() // Generate a new UUID
	return
}
