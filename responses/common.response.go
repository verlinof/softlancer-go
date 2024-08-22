package responses

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"` // Data can be of any type
}

type ErrorResponse struct {
	//* atau pointer digunakan agar nanti nilainya dapat berupa "nil"
	StatusCode int         `json:"status_code"`
	Error      interface{} `json:"error"`
}
