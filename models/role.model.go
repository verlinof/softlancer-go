package models

type Role struct {
	ID       *uint64 `json:"id" gorm:"primaryKey"`
	RoleName string  `json:"role_name" gorm:"type:varchar(255);not null;unique"`
}
