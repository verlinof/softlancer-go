package validations

import (
	"errors"
	"strings"

	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/models"
	"github.com/verlinof/softlancer-go/internal/requests"
)

func ValidateCreateApplication(request *requests.CreateApplicationRequest) error {
	var validationErrors []string

	if request.ProjectId == "" {
		validationErrors = append(validationErrors, "Project ID is required")
	} else {
		project := new(models.Project)
		database.DB.Table("projects").Where("id = ?", request.ProjectId).First(&project)
		if project.ID == "" {
			validationErrors = append(validationErrors, "Invalid project ID")
		}
	}

	if request.CuriculumVitae == "" {
		validationErrors = append(validationErrors, "Curiculum Vitae is required")
	}

	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, "; "))
	}

	return nil
}

func ValidateUpdateStatusApplication(request *requests.UpdateApplicationStatusRequest) error {
	var validationErrors []string
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

	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, "; "))
	}

	return nil
}
