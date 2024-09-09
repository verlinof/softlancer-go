package seeders

import (
	"fmt"
	"time"

	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/models"
)

func SeedRoles() error {
	roles := []models.Role{
		{
			RoleName: "Backend Developer",
		},
		{
			RoleName: "Frontend Developer",
		},
		{
			RoleName: "DevOps Engineer",
		},
		{
			RoleName: "Fullstack Developer",
		},
		{
			RoleName: "Mobile Developer",
		},
		{
			RoleName: "UI/UX Designer",
		},
		{
			RoleName: "Data Scientist",
		},
		{
			RoleName: "Data Analyst",
		},
	}

	var errMessages []string

	for _, role := range roles {
		// Create or find the role
		result := database.DB.Create(&role)
		if result.Error != nil {
			errMessages = append(errMessages, fmt.Sprintf("failed to seed role %s: %v", role.RoleName, result.Error))
		}
	}

	// Return error if any error messages exist
	if len(errMessages) > 0 {
		return fmt.Errorf("seeding errors: %v", errMessages)
	}

	time.Sleep(1 * time.Second)

	return nil
}
