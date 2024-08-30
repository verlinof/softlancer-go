package requests

type CreateApplicationRequest struct {
	UserId         uint64  `form:"user_id" json:"user_id"`
	ProjectId      *uint64 `form:"project_id" json:"project_id"`
	CuriculumVitae string  `form:"curiculum_vitae" json:"curiculum_vitae"`
	Portofolio     string  `form:"portofolio" json:"portofolio"`
	Status         string  `form:"status" json:"status"`
}

type UpdateApplicationRequest struct {
	CuriculumVitae string `form:"curiculum_vitae" json:"curiculum_vitae"`
	Portofolio     string `form:"portofolio" json:"portofolio"`
}
