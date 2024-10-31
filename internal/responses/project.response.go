package responses

type ProjectResponse struct {
	ID                 string `json:"id"`
	ProjectTitle       string `json:"project_title"`
	CompanyID          string `json:"company_id"`
	RoleID             string `json:"role_id"`
	ProjectDescription string `json:"project_description"`
	JobType            string `json:"job_type"`
	Status             string `json:"status"`
}
