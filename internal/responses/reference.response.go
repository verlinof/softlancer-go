package responses

type ReferenceResponse struct {
	ID     string `json:"id" gorm:"primaryKey"`
	UserID string `json:"user_id"`
	RoleID string `json:"role_id"`
}
