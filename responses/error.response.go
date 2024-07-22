package responses

type ErrorResponse struct {
	//* atau pointer digunakan agar nanti nilainya dapat berupa "nil"
	Status     string `json:"id"`
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}
