package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Reference struct {
	ID     string `json:"id" gorm:"type:varchar(36);primaryKey"`
	UserID string
	User   User `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RoleID string
	Role   Role `json:"role" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (reference *Reference) BeforeCreate(tx *gorm.DB) (err error) {
	reference.ID = uuid.New().String() // Generate a new UUID
	return
}
