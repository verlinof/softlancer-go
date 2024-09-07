package services

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"sync"
	"time"

	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/models"
)

// Struct untuk konfigurasi email
type emailConfig struct {
	SMTPHost    string
	SMTPPort    string
	SenderEmail string
	Password    string
	NumWorkers  int
	Auth        smtp.Auth
}

// emailService menggunakan emailConfig untuk konfigurasi
type EmailService struct {
	Config emailConfig
}

// Fungsi untuk membuat instance EmailService dengan konfigurasi yang diinisialisasi di awal
func NewEmailService() *EmailService {
	return &EmailService{
		Config: emailConfig{
			SMTPHost:    os.Getenv("MAIL_HOST"),
			SMTPPort:    os.Getenv("MAIL_PORT"),
			SenderEmail: os.Getenv("MAIL_USERNAME"),
			Password:    os.Getenv("MAIL_PASSWORD"),
			NumWorkers:  3, // atau bisa ditetapkan secara dinamis
			Auth:        smtp.PlainAuth("", os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_HOST")),
		},
	}
}

// Fungsi untuk mengirim email
func (e *EmailService) sendEmailJob(job models.EmailJob) {
	// Membuat header email termasuk subject
	headers := make(map[string]string)
	headers["From"] = e.Config.SenderEmail
	headers["To"] = job.To
	headers["Subject"] = job.Subject

	// Membuat format email dengan header dan body
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + job.Message

	// Otorisasi SMTP
	// auth := smtp.PlainAuth("", e.Config.SenderEmail, e.Config.Password, e.Config.SMTPHost)

	// Kirim email
	err := smtp.SendMail(e.Config.SMTPHost+":"+e.Config.SMTPPort, e.Config.Auth, e.Config.SenderEmail, []string{job.To}, []byte(message))
	if err != nil {
		log.Printf("Failed to send email to %s: %v", job.To, err)
		return
	}
	log.Printf("Email sent to %s successfully", job.To)
}

// Worker untuk memproses email secara paralel
func (e *EmailService) emailWorker(id int, jobs <-chan models.EmailJob, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		log.Printf("Worker %d: Sending email to %s", id, job.To)
		e.sendEmailJob(job)
		time.Sleep(3 * time.Second) // Simulasi waktu pemrosesan
	}
}

// Fungsi untuk mengirim email berdasarkan project
func (e *EmailService) SendEmail(roleId string, projectTitle string) (err error) {
	// Daftar email yang akan dikirim
	var emailList []map[string]interface{}
	var emailData []models.EmailJob
	// WaitGroup untuk menunggu semua emailWorker selesai
	var wg sync.WaitGroup

	err = database.DB.Table("users").
		Select(`
			users.id,
			users.email, 
			roles.id as role_id
		`).
		Joins("JOIN `references` ON users.id = `references`.user_id").
		Joins("JOIN `roles` ON `references`.role_id = `roles`.id").
		Where("roles.id = ?", roleId).
		Find(&emailList).Error

	fmt.Printf("%v", emailList)

	if err != nil {
		return err
	}

	for _, email := range emailList {
		emailData = append(emailData, models.EmailJob{
			To:      email["email"].(string),
			Subject: "Softlancer New Project",
			Message: fmt.Sprintf("New project: %s", projectTitle),
		})
	}

	// Channel untuk menampung email job
	jobs := make(chan models.EmailJob, len(emailList))

	// Mulai emailWorker pool
	for i := 1; i <= e.Config.NumWorkers; i++ {
		wg.Add(1)
		go e.emailWorker(i, jobs, &wg)
	}

	// Masukkan email ke dalam job queue
	for _, job := range emailData {
		jobs <- job
	}
	// Tutup channel jobs (semua job sudah dimasukkan)
	close(jobs)
	// Tunggu semua emailWorker selesai
	wg.Wait()

	log.Printf("Email sent to %d users successfully", len(emailList))
	return nil
}
