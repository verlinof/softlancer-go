package validations

import (
	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/models"
	"github.com/verlinof/softlancer-go/internal/requests"
)

func ValidateCreateCompany(request *requests.CreateCompanyRequest) []string {
	var validationErrors []string
	var companyExisted *models.Company

	//Validate Company Name
	if request.CompanyName == "" {
		validationErrors = append(validationErrors, "Company name is required")
	} else {
		database.DB.First(&companyExisted, "company_name = ?", request.CompanyName)
		if companyExisted.ID != "" {
			validationErrors = append(validationErrors, "Company name already exist")
		}
	}

	if request.CompanyDescription == "" {
		validationErrors = append(validationErrors, "Company description is required")
	}

	if request.CompanyLogo == nil {
		validationErrors = append(validationErrors, "Company logo is required")
	}

	return validationErrors
}

func ValidateUpdateCompany(request *requests.CreateCompanyRequest) []string {
	var validationErrors []string
	var companyExisted models.Company

	//Validate Company Name
	if request.CompanyName == "" {
		validationErrors = append(validationErrors, "Company name is required")
	} else {
		database.DB.First(&companyExisted, "company_name = ?", request.CompanyName)
		if companyExisted.ID != "" && request.CompanyName != companyExisted.CompanyName {
			validationErrors = append(validationErrors, "Company name already exist")
		}
	}

	if request.CompanyDescription == "" {
		validationErrors = append(validationErrors, "Company description is required")
	}

	if request.CompanyLogo == nil {
		validationErrors = append(validationErrors, "Company logo is required")
	}

	return validationErrors
}