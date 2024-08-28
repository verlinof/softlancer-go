package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/verlinof/softlancer-go/database"
	"github.com/verlinof/softlancer-go/models"
	"github.com/verlinof/softlancer-go/requests"
	"github.com/verlinof/softlancer-go/responses"
	"github.com/verlinof/softlancer-go/utils"
	"github.com/verlinof/softlancer-go/validations"
)

type CompanyController struct{}

func (e *CompanyController) Index(c *gin.Context) {
	var companyRes []responses.CompanyResponse
	err := database.DB.Table("companies").
		Select("id, company_name, company_description, company_logo").
		Scan(&companyRes).Error
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	message := "Success"
	if len(companyRes) == 0 {
		message = "Company data is empty"
	}

	successRes := responses.SuccessResponse{
		Message: message,
		Data:    companyRes,
	}

	c.JSON(http.StatusOK, successRes)
}

func (e *CompanyController) Show(c *gin.Context) {
	var err error
	var companyRes responses.CompanyResponse

	id := c.Param("id")
	//Error Handling
	parsedId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Invalid ID",
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	err = database.DB.Table("companies").
		Select("id, company_name, company_description, company_logo").
		Where("id = ?", parsedId).
		Scan(&companyRes).Error
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	if companyRes.ID == 0 {
		errResponse := responses.ErrorResponse{
			StatusCode: 404,
			Error:      "Company not found",
		}
		c.JSON(http.StatusNotFound, errResponse)
		return
	}

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
	validationErr := validations.ValidateCreateCompany(&companyReq)
	if len(validationErr) > 0 {
		errResponse := responses.ErrorResponse{
			StatusCode: 400,
			Error:      validationErr,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	// Upload File
	fileName, err := utils.HandleUploadFile(c, "company_logo", []string{"image/png", "image/jpeg", "image/jpg"}, 10000)
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	company := models.Company{
		CompanyName:        &companyReq.CompanyName,
		CompanyDescription: &companyReq.CompanyDescription,
		CompanyLogo:        &fileName,
	}

	// Create Company
	err = database.DB.Table("companies").Create(&company).Error
	if err != nil {
		utils.HandleRemoveFile(&fileName)
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Internal server error",
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

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
	parsedId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Invalid ID",
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	//Find the old data
	database.DB.Table("companies").Where("id = ?", parsedId).First(&oldCompany)

	if err = c.ShouldBind(&companyReq); err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 400,
			Error:      err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	// Validations
	validationErr := validations.ValidateUpdateCompany(&companyReq)
	if len(validationErr) > 0 {
		errResponse := responses.ErrorResponse{
			StatusCode: 400,
			Error:      validationErr,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	// Upload File
	fileName, err := utils.HandleUploadFile(c, "company_logo", []string{"image/png", "image/jpeg", "image/jpg"}, 10000)
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	company = models.Company{
		ID:                 &parsedId,
		CompanyName:        &companyReq.CompanyName,
		CompanyDescription: &companyReq.CompanyDescription,
		CompanyLogo:        &fileName,
	}

	// Update Company
	err = database.DB.Table("companies").Where("id = ?", parsedId).Updates(&company).Error
	if err != nil {
		utils.HandleRemoveFile(&fileName)
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Internal server error",
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	defer utils.HandleRemoveFile(oldCompany.CompanyLogo)

	successRes := responses.SuccessResponse{
		Message: "Success",
		Data:    company,
	}

	c.JSON(http.StatusOK, successRes)
}

func (e *CompanyController) Destroy(c *gin.Context) {
	id := c.Param("id")
	parsedId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Invalid ID",
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// Find the old data
	var company models.Company
	err = database.DB.Table("companies").Where("id = ?", parsedId).First(&company).Error
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
	if company.CompanyLogo != nil {
		defer utils.HandleRemoveFile(company.CompanyLogo)
	}

	// Delete the Company
	err = database.DB.Table("companies").Where("id = ?", parsedId).Delete(&company).Error
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
