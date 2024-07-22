package responses

type UserResponse struct {
	//* atau pointer digunakan agar nanti nilainya dapat berupa "nil"
	Id      *int    `json:"id" gorm:"primaryKey"`
	Name    *string `json:"name" gorm:"not null"`
	Address *string `json:"address" gorm:"not null"`
	Email   *string `json:"email" gorm:"not null"`
}

type LoginResponse struct {
	Status  string      `json:"id"`
	Message string      `json:"message"`
	Token   interface{} `json:"token"` // Data can be of any type
}
