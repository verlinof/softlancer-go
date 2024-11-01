package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/models"
	"github.com/verlinof/softlancer-go/internal/requests"
	"github.com/verlinof/softlancer-go/internal/responses"
	"github.com/verlinof/softlancer-go/internal/validations"
	"github.com/verlinof/softlancer-go/pkg"
)

type CompanyController struct{}

func (e *CompanyController) Index(c *gin.Context) {
	var companyRes []responses.CompanyResponse

	// Query the database
	err := database.DB.Table("companies").
		Select("id, company_name, company_description, company_logo").
		Scan(&companyRes).Error
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Failed to retrieve company data",
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// Update company logo URLs with the base API endpoint
	for i := range companyRes {
		logoPath := pkg.PrefixBaseUrl(companyRes[i].CompanyLogo)
		companyRes[i].CompanyLogo = *logoPath
	}

	// Prepare the response message
	message := "Success"
	if len(companyRes) == 0 {
		message = "No company data available"
	}

	successRes := responses.SuccessResponse{
		Message: message,
		Data:    companyRes,
	}

	// Return the JSON response
	c.JSON(http.StatusOK, successRes)
}

func (e *CompanyController) Show(c *gin.Context) {
	var err error
	var companyRes responses.CompanyResponse

	id := c.Param("id")

	err = database.DB.Table("companies").
		Select("id, company_name, company_description, company_logo").
		Where("id = ?", id).
		First(&companyRes).Error

	if err != nil && strings.Contains(err.Error(), "record not found") {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Company not found",
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	//Prefixing Base URL
	logoPath := pkg.PrefixBaseUrl(companyRes.CompanyLogo)
	companyRes.CompanyLogo = *logoPath

	successRes := responses.SuccessResponse{
		Message: "Success",
		Data:    companyRes,
	}

	c.JSON(http.StatusOK, successRes)
}

func (e *CompanyController) Store(c *gin.Context) {
	var err error
	var companyReq requests.CreateCompanyRequest

	if err = c.ShouldBind(&companyReq); err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 400,
			Error:      err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	// Validations
	err = validations.ValidateCreateCompany(&companyReq)
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 400,
			Error:      err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	// Upload File
	fileName, err := pkg.HandleUploadFile(c, "companies", "company_logo", []string{"image/png", "image/jpeg", "image/jpg"}, 10000)
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	company := models.Company{
		CompanyName:        companyReq.CompanyName,
		CompanyDescription: companyReq.CompanyDescription,
		CompanyLogo:        fileName,
	}

	// Create Company
	err = database.DB.Table("companies").Create(&company).Error
	if err != nil {
		pkg.HandleRemoveFile(fileName)
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Internal server error",
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	company.CompanyLogo = *pkg.PrefixBaseUrl(company.CompanyLogo)

	successRes := responses.SuccessResponse{
		Message: "Success",
		Data:    company,
	}

	c.JSON(http.StatusOK, successRes)
}

func (e *CompanyController) Update(c *gin.Context) {
	var err error
	var companyReq requests.CreateCompanyRequest
	var company models.Company
	var oldCompany models.Company

	id := c.Param("id")

	//Find the old data
	err = database.DB.Table("companies").Where("id = ?", id).First(&oldCompany).Error
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 404,
			Error:      "Company not found",
		}
		c.JSON(http.StatusNotFound, errResponse)
		return
	}

	if err = c.ShouldBind(&companyReq); err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 400,
			Error:      err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	// Validations
	err = validations.ValidateUpdateCompany(&companyReq)
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 400,
			Error:      err,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	// Upload File
	fileName, err := pkg.HandleUploadFile(c, "companies", "company_logo", []string{"image/png", "image/jpeg", "image/jpg"}, 10000)
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	company = models.Company{
		ID:                 id,
		CompanyName:        companyReq.CompanyName,
		CompanyDescription: companyReq.CompanyDescription,
		CompanyLogo:        fileName,
	}

	// Update Company
	err = database.DB.Table("companies").Where("id = ?", id).Updates(&company).Error
	if err != nil {
		pkg.HandleRemoveFile(fileName)
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Internal server error",
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	defer pkg.HandleRemoveFile(oldCompany.CompanyLogo)

	company.CompanyLogo = *pkg.PrefixBaseUrl(company.CompanyLogo)

	successRes := responses.SuccessResponse{
		Message: "Success",
		Data:    company,
	}

	c.JSON(http.StatusOK, successRes)
}

func (e *CompanyController) Destroy(c *gin.Context) {
	id := c.Param("id")

	// Find the old data
	var company models.Company
	err := database.DB.Table("companies").Where("id = ?", id).First(&company).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			errResponse := responses.ErrorResponse{
				StatusCode: 404,
				Error:      "Company not found",
			}
			c.JSON(http.StatusNotFound, errResponse)
			return
		}

		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Internal server error",
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// Remove the file if it exists
	if company.CompanyLogo != "" {
		defer pkg.HandleRemoveFile(company.CompanyLogo)
	}

	// Delete the Company
	err = database.DB.Table("companies").Where("id = ?", id).Delete(&company).Error
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Internal server error",
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	successRes := responses.SuccessResponse{
		Message: "Success",
		Data:    &company,
	}

	c.JSON(http.StatusOK, successRes)
}
