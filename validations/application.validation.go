package validations

import (
	"github.com/verlinof/softlancer-go/database"
	"github.com/verlinof/softlancer-go/models"
	"github.com/verlinof/softlancer-go/requests"
)

func ValidateCreateApplication(request *requests.CreateApplicationRequest) []string {
	var validationErrors []string

	if request.ProjectId == nil {
		validationErrors = append(validationErrors, "Project ID is required")
	} else {
		project := new(models.Project)
		database.DB.Table("projects").Where("id = ?", *request.ProjectId).First(&project)
		if project.ID == nil {
			validationErrors = append(validationErrors, "Invalid project ID")
		}
	}

	if request.CuriculumVitae == "" {
		validationErrors = append(validationErrors, "Curiculum Vitae is required")
	}

	validStatus := map[string]bool{
		"waiting":  true,
		"accepted": true,
		"rejected": true,
	}

	if request.Status == "" {
		validationErrors = append(validationErrors, "Status is required")
	} else {
		if !validStatus[request.Status] {
			validationErrors = append(validationErrors, "Invalid status. Allowed values are: waiting, accepted, rejected")
		}
	}

	return validationErrors
}
