package models

import "gorm.io/gorm"

type Reference struct {
	gorm.Model
	Id        *int    `json:"id" gorm:"primaryKey"`
	UserID    *int 
	User      *User   `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RoleID 		*int 
	Role 			*Role 	`json:"role" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}