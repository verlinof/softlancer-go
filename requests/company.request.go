package requests

import "mime/multipart"

type CreateCompanyRequest struct {
	CompanyName        string                `form:"company_name" json:"company_name"`
	CompanyDescription string                `form:"company_description" json:"company_description"`
	CompanyLogo        *multipart.FileHeader `form:"company_logo" json:"company_logo"`
}
