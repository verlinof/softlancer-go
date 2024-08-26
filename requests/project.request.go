package requests

type ProjectRequest struct {
	ProjectTitle       string `json:"project_title"`
	ProjectDescription string `json:"project_description"`
	CompanyId          *int   `json:"company_id"`
	RoleId             *int   `json:"role_id"`
	JobType            string `json:"job_type"`
	Status             string `json:"status"`
}
