package seeders

import (
	"fmt"

	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/models"
)

func SeedReferences() error {
	var users []models.User
	var roles []models.Role

	// Ambil semua users
	result := database.DB.Table("users").Find(&users)
	if result.Error != nil {
		return fmt.Errorf("failed to fetch users: %v", result.Error)
	}
	if len(users) == 0 {
		return fmt.Errorf("no users found, cannot seed references")
	}

	// Ambil semua roles
	result = database.DB.Table("roles").Find(&roles)
	if result.Error != nil {
		return fmt.Errorf("failed to fetch roles: %v", result.Error)
	}
	if len(roles) == 0 {
		return fmt.Errorf("no roles found, cannot seed references")
	}

	// Seed references
	references := []models.Reference{
		{UserID: users[0].ID, RoleID: roles[0].ID},
		{UserID: users[0].ID, RoleID: roles[1].ID},
		{UserID: users[0].ID, RoleID: roles[2].ID},
		{UserID: users[1].ID, RoleID: roles[3].ID},
		{UserID: users[1].ID, RoleID: roles[4].ID},
		{UserID: users[1].ID, RoleID: roles[5].ID},
	}

	var errMessages []string

	for _, reference := range references {
		// FirstOrCreate reference
		result := database.DB.Create(&reference)
		if result.Error != nil {
			errMessages = append(errMessages, fmt.Sprintf("Failed to seed reference for user %s and role %s: %v", reference.UserID, reference.RoleID, result.Error))
		}
	}

	// Kembalikan semua error jika ada
	if len(errMessages) > 0 {
		return fmt.Errorf("seeding errors: %v", errMessages)
	}

	return nil
}
