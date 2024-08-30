package models

type Application struct {
	ID             *uint64 `json:"id" gorm:"primaryKey"`
	UserID         *uint64
	User           User `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProjectID      *uint64
	Project        Project `json:"project" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CuriculumVitae string  `json:"curiculum_vitae" gorm:"type:varchar(255);not null"`
	Portofolio     string  `json:"portofolio" gorm:"type:varchar(255);not null"`
	Status         string  `json:"status" default:"waiting" gorm:"type:enum('waiting','accepted','rejected');not null"`
}
