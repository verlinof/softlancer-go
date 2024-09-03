package requests

type CreateApplicationRequest struct {
	UserId         string  `form:"user_id" json:"user_id"`
	ProjectId      *uint64 `form:"project_id" json:"project_id"`
	CuriculumVitae string  `form:"curiculum_vitae" json:"curiculum_vitae"`
	Portofolio     string  `form:"portofolio" json:"portofolio"`
}

type UpdateApplicationRequest struct {
	CuriculumVitae string `form:"curiculum_vitae" json:"curiculum_vitae"`
	Portofolio     string `form:"portofolio" json:"portofolio"`
}

type UpdateApplicationStatusRequest struct {
	Status string `form:"status" json:"status"`
}
