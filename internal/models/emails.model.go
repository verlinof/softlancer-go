package models

// Job struct untuk mengirim email dengan subject
type EmailJob struct {
	To      string
	Subject string
	Message string
}
