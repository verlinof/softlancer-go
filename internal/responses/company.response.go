package responses

type CompanyResponse struct {
	ID                 string `json:"id"`
	CompanyName        string `json:"company_name"`
	CompanyDescription string `json:"company_description"`
	CompanyLogo        string `json:"company_logo"`
}
