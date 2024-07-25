package responses

type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"` // Data can be of any type
}

type ErrorResponse struct {
	//* atau pointer digunakan agar nanti nilainya dapat berupa "nil"
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}
