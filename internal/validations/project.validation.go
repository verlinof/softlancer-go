package validations

import (
	"errors"
	"strings"

	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/models"
	"github.com/verlinof/softlancer-go/internal/requests"
)

// Validation
func ValidateCreateProject(request *requests.ProjectRequest) error {
	var validationErrors []string

	if request.ProjectTitle == "" {
		validationErrors = append(validationErrors, "Email is required")
	}

	if request.ProjectDescription == "" {
		validationErrors = append(validationErrors, "Email is required")
	}

	// =====Job Type=====
	if request.JobType == "" {
		validationErrors = append(validationErrors, "job type is required")
	}
	// Daftar nilai yang diizinkan untuk JobType
	validJobTypes := map[string]bool{
		"fulltime":  true,
		"parttime":  true,
		"freelance": true,
	}

	// Cek apakah JobType valid
	if !validJobTypes[request.JobType] {
		validationErrors = append(validationErrors, "Invalid job type. Allowed values are: fulltime, parttime, freelance")
	}

	validStatus := map[string]bool{
		"open":   true,
		"closed": true,
	}

	if request.Status != "" && !validStatus[request.Status] {
		validationErrors = append(validationErrors, "Invalid status. Allowed values are: open, closed")
	}

	// =====Company ID=====
	if request.CompanyId == "" {
		validationErrors = append(validationErrors, "Company ID is required")
	}

	if (request.CompanyId) != "" {
		company := new(models.Company)
		database.DB.Table("companies").Where("id = ?", request.CompanyId).First(&company)
		if company.ID == "" {
			validationErrors = append(validationErrors, "Invalid company ID")
		}
	}

	// =====Role ID=====
	if request.RoleId == "" {
		validationErrors = append(validationErrors, "Role id is required")
	}
	if request.RoleId != "" {
		role := new(models.Role)
		database.DB.Table("roles").Where("id = ?", request.RoleId).First(&role)
		if role.ID == "" {
			validationErrors = append(validationErrors, "Invalid role ID")
		}
	}

	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, "; "))
	}

	return nil
}

func ValidateUpdateProject(request *requests.ProjectRequest) error {
	var validationErrors []string
	// =====Job Type=====
	// Daftar nilai yang diizinkan untuk JobType
	validJobTypes := map[string]bool{
		"fulltime":  true,
		"parttime":  true,
		"freelance": true,
	}

	// Cek apakah JobType valid
	if request.JobType != "" && !validJobTypes[request.JobType] {
		validationErrors = append(validationErrors, "Invalid job type. Allowed values are: fulltime, parttime, freelance")
	}

	// =====Status=====
	validStatus := map[string]bool{
		"open":   true,
		"closed": true,
	}

	if request.Status != "" && !validStatus[request.Status] {
		validationErrors = append(validationErrors, "Invalid status. Allowed values are: open, closed")
	}

	// =====Company ID=====
	if request.CompanyId != "" {
		company := new(models.Company)
		database.DB.Table("companies").Where("id = ?", request.CompanyId).First(&company)
		if company.ID == "" {
			validationErrors = append(validationErrors, "Invalid company ID")
		}
	}

	// =====Role ID=====
	if request.RoleId != "" {
		role := new(models.Role)
		database.DB.Table("roles").Where("id = ?", request.RoleId).First(&role)
		if role.ID == "" {
			validationErrors = append(validationErrors, "Invalid role ID")
		}
	}

	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, "; "))
	}

	return nil
}
