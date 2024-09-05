package responses

type ApplicationResponse struct {
	ID             string `json:"id"`
	ProjectID      string `json:"project_id"`
	ProjectTitle   string `json:"project_title"`
	RoleID         string `json:"role_id"`
	RoleName       string `json:"role_name"`
	CuriculumVitae string `json:"curiculum_vitae"`
	Portofolio     string `json:"portofolio"`
	Status         string `json:"status"`
}

type ApplicationStoreResponse struct {
	ID             string `json:"id"`
	UserID         string `json:"user_id"`
	ProjectID      string `json:"project_id"`
	CuriculumVitae string `json:"curiculum_vitae"`
	Portofolio     string `json:"portofolio"`
	Status         string `json:"status"`
}

type ApplicationDetailResponse struct {
	ID                 string `json:"id"`
	ProjectID          string `json:"project_id"`
	ProjectTitle       string `json:"project_title"`
	ProjectDescription string `json:"project_description"`
	UserID             string `json:"user_id"`
	UserName           string `json:"user_name"`
	UserEmail          string `json:"user_email"`
	CompanyID          string `json:"company_id"`
	CompanyName        string `json:"company_name"`
	CompanyDescription string `json:"company_description"`
	CompanyLogo        string `json:"company_logo"`
	RoleID             string `json:"role_id"`
	RoleName           string `json:"role_name"`
	CuriculumVitae     string `json:"curiculum_vitae"`
	Portofolio         string `json:"portofolio"`
	Status             string `json:"status"`
}
