package models

type Company struct {
	ID                 *uint64 `json:"id" gorm:"primaryKey"`
	CompanyName        string  `json:"company_name" gorm:"type:varchar(255);not null;unique"`
	CompanyDescription string  `json:"company_description" gorm:"type:varchar(255);not null"`
	CompanyLogo        string  `json:"company_logo" gorm:"type:varchar(255);not null"`
}
