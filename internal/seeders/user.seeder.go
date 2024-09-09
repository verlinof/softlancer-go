package seeders

import (
	"fmt"
	"time"

	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func SeedUsers() error {
	users := []models.User{
		{
			Name:     "Verlino Fajri",
			Email:    "marimo.zx@gmail.com",
			Address:  "Jl. Pemuda No. 1",
			Password: "fajri123",
			IsAdmin:  true,
		},
		{
			Name:     "Verlinof",
			Email:    "ajikoko.zx@gmail.com",
			Address:  "Jl. Pemuda No. 1",
			Password: "fajri123",
			IsAdmin:  false,
		},
	}

	var errMessages []string

	for _, user := range users {
		// Encrypt the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			errMessages = append(errMessages, fmt.Sprintf("failed to encrypt password for user %s: %v", user.Email, err))
			continue
		}
		user.Password = string(hashedPassword)

		// Try to create or find the user
		result := database.DB.Create(&user)
		if result.Error != nil {
			errMessages = append(errMessages, fmt.Sprintf("failed to seed user %s: %v", user.Email, result.Error))
		}
	}

	// Return error if any error messages exist
	if len(errMessages) > 0 {
		return fmt.Errorf("seeding errors: %v", errMessages)
	}

	time.Sleep(1 * time.Second)

	return nil
}
