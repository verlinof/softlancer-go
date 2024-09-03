package models

type Reference struct {
	ID     *uint `json:"id" gorm:"primaryKey"`
	UserID *uint64
	User   User `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RoleID uint
	Role   Role `json:"role" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
