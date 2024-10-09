package validations

import (
	"errors"
	"regexp"
	"strings"

	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/models"
	"github.com/verlinof/softlancer-go/internal/requests"
)

// Validation
func ValidateLogin(request *requests.LoginRequest) error {
	var validationErrors []string

	// Validate email is not empty
	if request.Email == "" {
		validationErrors = append(validationErrors, "Email is required")
	}

	// Validate email format using regex
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if request.Email != "" && !emailRegex.MatchString(request.Email) {
		validationErrors = append(validationErrors, "Invalid email format")
	}

	// Validate password is not empty
	if request.Password == "" {
		validationErrors = append(validationErrors, "Password is required")
	}

	// If there are validation errors, return them as a single error
	if len(validationErrors) > 0 {
		// Join all validation errors into one error message
		return errors.New(strings.Join(validationErrors, ";"))
	}

	return nil
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
