package requests

import "time"

type UserRequest struct {
	Name      string    `json:"name" binding:"required"`
	Address   string    `json:"address" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Born_date time.Time `json:"born_date" binding:"required"`
	Password  string    `json:"password" binding:"required"`
} //Buat ngehandle validation