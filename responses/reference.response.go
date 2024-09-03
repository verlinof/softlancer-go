package responses

type ReferenceResponse struct {
	ID     *uint  `json:"id" gorm:"primaryKey"`
	UserID uint64 `json:"user_id"`
	RoleID uint   `json:"role_id"`
}
