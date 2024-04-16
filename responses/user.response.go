package responses

type UserResponse struct {
	//* atau pointer digunakan agar nanti nilainya dapat berupa "nil"
	Id      *int    `json:"id" gorm:"primaryKey"`
	Name    *string `json:"name" gorm:"not null"`
	Address *string `json:"address" gorm:"not null"`
	Email   *string `json:"email" gorm:"not null"`
}