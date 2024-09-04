package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID       string `json:"id" gorm:"type:varchar(36);primaryKey"`
	RoleName string `json:"role_name" gorm:"type:varchar(255);not null;unique"`
}

func (role *Role) BeforeCreate(tx *gorm.DB) (err error) {
	role.ID = uuid.New().String()
	return
}
