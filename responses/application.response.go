package responses

type ApplicationResponse struct {
	ID             *uint64 `json:"id"`
	ProjectID      *uint64 `json:"project_id"`
	ProjectTitle   string  `json:"project_title"`
	RoleID         *uint64 `json:"role_id"`
	RoleName       string  `json:"role_name"`
	CuriculumVitae string  `json:"curiculum_vitae"`
	Portofolio     string  `json:"portofolio"`
	Status         string  `json:"status"`
}

type ApplicationStoreResponse struct {
	ID             *uint64 `json:"id"`
	UserID         *uint64 `json:"user_id"`
	ProjectID      *uint64 `json:"project_id"`
	CuriculumVitae string  `json:"curiculum_vitae"`
	Portofolio     string  `json:"portofolio"`
	Status         string  `json:"status"`
}

type ApplicationDetailResponse struct {
	ID                 *uint64 `json:"id"`
	ProjectID          *uint64 `json:"project_id"`
	ProjectTitle       string  `json:"project_title"`
	ProjectDescription string  `json:"project_description"`
	CompanyID          *uint64 `json:"company_id"`
	CompanyName        string  `json:"company_name"`
	CompanyDescription string  `json:"company_description"`
	CompanyLogo        string  `json:"company_logo"`
	RoleID             *uint64 `json:"role_id"`
	RoleName           string  `json:"role_name"`
	CuriculumVitae     string  `json:"curiculum_vitae"`
	Portofolio         string  `json:"portofolio"`
	Status             string  `json:"status"`
}
