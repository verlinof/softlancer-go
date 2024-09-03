package validations

import (
	"regexp"

	"github.com/verlinof/softlancer-go/database"
	"github.com/verlinof/softlancer-go/models"
	"github.com/verlinof/softlancer-go/requests"
)

// Validation
func ValidateLogin(request *requests.LoginRequest) []string {
	var validationErrors []string

	if request.Email == "" {
		validationErrors = append(validationErrors, "Email is required")
	}

	// Regex untuk validasi format email
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(request.Email) {
		validationErrors = append(validationErrors, "Invalid email format")
	}

	if request.Password == "" {
		validationErrors = append(validationErrors, "Password is required")
	}

	return validationErrors
}

func ValidateRegister(request *requests.UserRequest) []string {
	var validationErrors []string

	//Check if the email already exist
	userEmailExist := new(models.User)
	database.DB.Table("users").Where("email = ?", request.Email).First(&userEmailExist)
	if userEmailExist.ID != "" {
		validationErrors = append(validationErrors, "Email already exist")
		return validationErrors
	}
	if request.Email == "" {
		validationErrors = append(validationErrors, "Email is required")
	}
	// Regex untuk validasi format email
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(request.Email) {
		validationErrors = append(validationErrors, "Invalid email format")
	}

	if request.Name == "" {
		validationErrors = append(validationErrors, "Name is requiredd")
	}

	if request.Address == "" {
		validationErrors = append(validationErrors, "Address is required")
	}

	if request.Password == "" {
		validationErrors = append(validationErrors, "Password is required")
	}

	return validationErrors
}
