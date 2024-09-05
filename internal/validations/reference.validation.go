package validations

import (
	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/models"
	"github.com/verlinof/softlancer-go/internal/requests"
)

func ValidateCreateReference(request *requests.CreateReferenceRequest) []string {
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
	return validationErrors
}
