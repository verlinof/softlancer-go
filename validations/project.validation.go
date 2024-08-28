package validations

import (
	"github.com/verlinof/softlancer-go/database"
	"github.com/verlinof/softlancer-go/models"
	"github.com/verlinof/softlancer-go/requests"
)

// Validation
func ValidateCreateProject(request *requests.ProjectRequest) []string {
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

	// =====Status=====
	if request.Status == "" {
		validationErrors = append(validationErrors, "Status is required")
	}

	validStatus := map[string]bool{
		"open":   true,
		"closed": true,
	}

	if !validStatus[request.Status] {
		validationErrors = append(validationErrors, "Invalid status. Allowed values are: open, closed")
	}

	// =====Company ID=====
	if request.CompanyId == 0 {
		validationErrors = append(validationErrors, "Company ID is required")
	}

	if (request.CompanyId) != 0 {
		company := new(models.Company)
		database.DB.Table("companies").Where("id = ?", request.CompanyId).First(&company)
		if company.ID == nil {
			validationErrors = append(validationErrors, "Invalid company ID")
		}
	}

	// =====Role ID=====
	if request.RoleId == 0 {
		validationErrors = append(validationErrors, "Role id is required")
	}
	if request.RoleId != 0 {
		role := new(models.Role)
		database.DB.Table("roles").Where("id = ?", request.RoleId).First(&role)
		if role.ID == nil {
			validationErrors = append(validationErrors, "Invalid role ID")
		}
	}

	return validationErrors
}

func ValidateUpdateProject(request *requests.ProjectRequest) []string {
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
	if (request.CompanyId) != 0 {
		company := new(models.Company)
		database.DB.Table("companies").Where("id = ?", request.CompanyId).First(&company)
		if company.ID == nil {
			validationErrors = append(validationErrors, "Invalid company ID")
		}
	}

	// =====Role ID=====
	if request.RoleId != 0 {
		role := new(models.Role)
		database.DB.Table("roles").Where("id = ?", request.RoleId).First(&role)
		if role.ID == nil {
			validationErrors = append(validationErrors, "Invalid role ID")
		}
	}

	return validationErrors
}
