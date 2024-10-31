package validations

import (
	"errors"
	"strings"

	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/models"
	"github.com/verlinof/softlancer-go/internal/requests"
)

func ValidateCreateReference(request *requests.CreateReferenceRequest) error {
	var validationErrors []string

	if request.RoleID == "" {
		validationErrors = append(validationErrors, "Project ID is required")
	} else {
		role := new(models.Role)
		database.DB.Table("roles").Where("id = ?", request.RoleID).First(&role)
		if role.ID == "" {
			validationErrors = append(validationErrors, "Invalid project ID")
		}
	}

	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, "; "))
	}

	return nil
}
