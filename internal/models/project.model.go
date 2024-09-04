package models

type Project struct {
	ID                 *uint64 `json:"id" gorm:"primaryKey"`
	CompanyID          string
	Company            Company `json:"company" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RoleID             string
	Role               Role   `json:"role" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProjectTitle       string `json:"project_title" gorm:"type:varchar(255);not null"`
	ProjectDescription string `json:"project_description" gorm:"type:text;not null"`
	JobType            string `json:"job_type" gorm:"type:enum('fulltime','parttime','freelance');not null"`
	Status             string `json:"status" gorm:"type:enum('open','closed');not null"`
}
