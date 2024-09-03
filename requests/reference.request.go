package requests

type CreateReferenceRequest struct {
	RoleID *uint `form:"role_id" json:"role_id"`
}
