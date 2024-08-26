package models

type Role struct {
	ID       *int    `json:"id" gorm:"primaryKey"`
	RoleName *string `json:"role_name" gorm:"type:varchar(255);not null"`
}
