package services

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"sync"

	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/models"
)

// emailService menggunakan emailConfig untuk konfigurasi
type EmailService struct {
	SMTPHost    string
	SMTPPort    string
	MailName    string
	SenderEmail string
	Password    string
	NumWorkers  int
	Auth        smtp.Auth
}

// Fungsi untuk membuat instance EmailService dengan konfigurasi yang diinisialisasi di awal
func NewEmailService() *EmailService {
	return &EmailService{
		SMTPHost:    os.Getenv("MAIL_HOST"),
		SMTPPort:    os.Getenv("MAIL_PORT"),
		MailName:    os.Getenv("MAIL_FROM_NAME"),
		SenderEmail: os.Getenv("MAIL_USERNAME"),
		Password:    os.Getenv("MAIL_PASSWORD"),
		NumWorkers:  3, // atau bisa ditetapkan secara dinamis
		Auth:        smtp.PlainAuth("", os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_HOST")),
	}
}

// Fungsi untuk mengirim email berdasarkan project
func (e *EmailService) SendEmail(roleId string, projectTitle string) {
	// Daftar email yang akan dikirim
	var emailList []map[string]interface{}
	var emailData []models.EmailJob
	// WaitGroup untuk menunggu semua emailWorker selesai
	var wg sync.WaitGroup

	err := database.DB.Table("users").
		Select(`
			users.id,
			users.email, 
			roles.id as role_id
		`).
		Joins("JOIN `references` ON users.id = `references`.user_id").
		Joins("JOIN `roles` ON `references`.role_id = `roles`.id").
		Where("roles.id = ?", roleId).
		Find(&emailList).Error

	if err != nil {
		log.Printf("Failed to get email list: %v", err)
	}

	for _, email := range emailList {
		message := fmt.Sprintf(`
Job Sesuai Referensimu Telah Tersedia
Hi %s,

Kami berharap Anda dalam keadaan baik. Kami senang memberi tahu Anda bahwa ada pekerjaan baru yang sesuai dengan preferensi Anda yang telah tersedia di aplikasi SoftLancer!
Kami mengundang Anda untuk segera memeriksa detail pekerjaan tersebut. Jangan lewatkan kesempatan ini untuk mengembangkan keterampilan Anda dan mendapatkan pengalaman berharga.

Terima kasih telah menjadi bagian dari SoftLancer. Kami berharap dapat membantu Anda dalam perjalanan karir Anda.

Salam Hormat,
Tim Softlancer
    `, email["email"])

		emailData = append(emailData, models.EmailJob{
			To:      email["email"].(string),
			Subject: "Softlancer New Project",
			Message: message,
		})
	}

	// Channel untuk menampung email job
	jobs := make(chan models.EmailJob, len(emailList))
	for i := 1; i <= e.NumWorkers; i++ {
		wg.Add(1)
		go e.emailWorker(i, jobs, &wg)
	}
	for _, job := range emailData {
		jobs <- job
	}

	// Tutup channel jobs (semua job sudah dimasukkan)
	close(jobs)
	// Tunggu semua emailWorker selesai
	wg.Wait()

	fmt.Println(e.SMTPHost, e.SMTPPort, e.MailName, e.SenderEmail, e.Password, e.NumWorkers, e.Auth)
}

// Fungsi untuk mengirim email
func (e *EmailService) sendEmailJob(job models.EmailJob) {
	// Membuat header email termasuk subject
	headers := make(map[string]string)
	headers["From"] = e.MailName
	headers["To"] = job.To
	headers["Subject"] = job.Subject

	// Membuat format email dengan header dan body
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + job.Message

	// Kirim email
	err := smtp.SendMail(e.SMTPHost+":"+e.SMTPPort, e.Auth, e.SenderEmail, []string{job.To}, []byte(message))
	if err != nil {
		log.Printf("Failed to send email to %s: %v", job.To, err)
		return
	}
}

// Worker untuk memproses email secara paralel
func (e *EmailService) emailWorker(id int, jobs <-chan models.EmailJob, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		log.Printf("Worker %d: Sending email to %s", id, job.To)
		e.sendEmailJob(job)
	}
}
