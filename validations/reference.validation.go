package validations

import (
	"github.com/verlinof/softlancer-go/database"
	"github.com/verlinof/softlancer-go/models"
	"github.com/verlinof/softlancer-go/requests"
)

func ValidateCreateReference(request *requests.CreateReferenceRequest) []string {
	var validationErrors []string

	if request.RoleID == nil {
		validationErrors = append(validationErrors, "Project ID is required")
	} else {
		project := new(models.Project)
		database.DB.Table("projects").Where("id = ?", *request.RoleID).First(&project)
		if project.ID == nil {
			validationErrors = append(validationErrors, "Invalid project ID")
		}
	}
	return validationErrors
}
