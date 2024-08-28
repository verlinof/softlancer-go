package responses

type ProjectResponse struct {
	ID                 *uint64 `json:"id"`
	ProjectTitle       string  `json:"project_title"`
	ProjectDescription string  `json:"project_description"`
	JobType            string  `json:"job_type"`
	Status             string  `json:"status"`
}
