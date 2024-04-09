package models

import (
	"time"
)

type User struct {
	// gorm.Model
	Id        int     `json:"id" gorm:"primaryKey"`
	Name      string  `json:"name" gorm:"not null"`
	Address   string  `json:"address" gorm:"not null"`
	Email     string  `json:"email" gorm:"not null"`
	Password  string  `json:"password" gorm:"not null"`
	Born_date time.Time `json:"born_date"`
}