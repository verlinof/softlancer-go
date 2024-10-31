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

type ApplicationController struct{}

func (a *ApplicationController) Index(c *gin.Context) {
	var response []responses.ApplicationResponse
	err := database.DB.Table("applications").
		Joins("JOIN projects ON applications.project_id = projects.id").
		Joins("JOIN roles ON projects.role_id = roles.id").
		Select(`
		applications.id,
		projects.id as project_id,
		projects.project_title, 
		roles.id as role_id,
		roles.role_name,
		applications.curiculum_vitae,
		applications.portofolio,
		applications.status
		`).
		Scan(&response).Error

	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	message := "Success"
	if len(response) == 0 {
		message = "Applications data is empty"
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: message,
		Data:    response,
	})
}

func (e *ApplicationController) Show(c *gin.Context) {
	var response responses.ApplicationDetailResponse

	id := c.Param("id")
	err := database.DB.Table("applications").
		Joins("Join users ON applications.user_id = users.id").
		Joins("JOIN projects ON applications.project_id = projects.id").
		Joins("JOIN companies ON projects.company_id = companies.id").
		Joins("JOIN roles ON projects.role_id = roles.id").
		Select(`
			applications.id,
			projects.id as project_id,
			projects.project_title, 
			projects.project_description,
			users.id as user_id,
			users.name as user_name,
			users.email as user_email,
			companies.id as company_id,
			companies.company_name,
			companies.company_description,
			companies.company_logo,
			roles.id as role_id,
			roles.role_name,
			applications.curiculum_vitae,
			applications.portofolio,
			applications.status
		`).
		Where("applications.id = ?", id).
		First(&response).Error

	response.CompanyLogo = *pkg.PrefixBaseUrl(response.CompanyLogo)

	if err != nil && strings.Contains(err.Error(), "not found") {
		errResponse := responses.ErrorResponse{
			StatusCode: 404,
			Error:      "Application not found",
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Success",
		Data:    response,
	})
}

func (e *ApplicationController) Store(c *gin.Context) {
	var err error
	var request requests.CreateApplicationRequest
	var application models.Application

	//Get User id from Middleware
	userId, _ := c.Get("user")

	// Bind request body ke struct ProjectRequest
	if err = c.ShouldBind(&request); err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 400,
			Error:      "Invalid request body",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	// Validasi input dari user
	err = validations.ValidateCreateApplication(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ErrorResponse{
			StatusCode: 400,
			Error:      err.Error(),
		})
		return
	}

	// Mengisi model application berdasarkan request
	application = models.Application{
		UserID:         userId.(string),
		ProjectID:      request.ProjectId,
		CuriculumVitae: request.CuriculumVitae,
		Portofolio:     request.Portofolio,
	}

	// Simpan project ke database
	if err = database.DB.Create(&application).Error; err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// Mengembalikan response sukses
	successRes := responses.SuccessResponse{
		Message: "Success to create application",
		Data: responses.ApplicationStoreResponse{
			ID:             application.ID,
			UserID:         application.UserID,
			ProjectID:      application.ProjectID,
			CuriculumVitae: application.CuriculumVitae,
			Portofolio:     application.Portofolio,
			Status:         application.Status,
		},
	}

	c.JSON(http.StatusCreated, successRes)
}

func (e *ApplicationController) Update(c *gin.Context) {
	var err error
	var request requests.UpdateApplicationRequest
	var application models.Application

	id := c.GetString("application")

	//Find the old data
	err = database.DB.Table("applications").Where("id = ?", id).First(&application).Error
	if err != nil && strings.Contains(err.Error(), "record not found") {
		errResponse := responses.ErrorResponse{
			StatusCode: 404,
			Error:      "Application not found",
		}
		c.JSON(http.StatusNotFound, errResponse)
		return
	}

	// Bind request body ke struct ProjectRequest
	if err = c.ShouldBind(&request); err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 400,
			Error:      "Invalid request body",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	// Simpan application ke database
	if err = database.DB.Table("applications").Where("id = ?", id).Updates(&request).Error; err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// Mencari data yang telah diupdate
	database.DB.Table("applications").Where("id = ?", id).First(&application)

	// Mengembalikan response sukses
	successRes := responses.SuccessResponse{
		Message: "Success Update Application",
		Data: responses.ApplicationStoreResponse{
			ID:             application.ID,
			UserID:         application.UserID,
			ProjectID:      application.ProjectID,
			CuriculumVitae: application.CuriculumVitae,
			Portofolio:     application.Portofolio,
			Status:         application.Status,
		},
	}

	c.JSON(http.StatusOK, successRes)
}

func (e *ApplicationController) UpdateStatus(c *gin.Context) {
	var err error
	var request requests.UpdateApplicationStatusRequest
	var application models.Application

	applicationId := c.Param("id")

	//Find the application
	err = database.DB.Table("applications").Where("id = ?", applicationId).First(&application).Error
	if err != nil && strings.Contains(err.Error(), "record not found") {
		errResponse := responses.ErrorResponse{
			StatusCode: 404,
			Error:      "Application not found",
		}
		c.AbortWithStatusJSON(http.StatusNotFound, errResponse)
		return
	}

	// Bind request body ke struct ProjectRequest
	if err = c.ShouldBind(&request); err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 400,
			Error:      "Invalid request body",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	// Simpan application ke database
	if err = database.DB.Table("applications").Where("id = ?", applicationId).Updates(&request).Error; err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// Mencari data yang telah diupdate
	database.DB.Table("applications").Where("id = ?", applicationId).First(&application)

	// Mengembalikan response sukses
	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Success Update Application",
		Data: responses.ApplicationStoreResponse{
			ID:             application.ID,
			UserID:         application.UserID,
			ProjectID:      application.ProjectID,
			CuriculumVitae: application.CuriculumVitae,
			Portofolio:     application.Portofolio,
			Status:         application.Status,
		},
	})
}

func (e *ApplicationController) Destroy(c *gin.Context) {
	var application models.Application

	id := c.GetString("application")

	//Find the old data
	database.DB.Table("applications").Where("id = ?", id).First(&application)

	err := database.DB.Table("applications").Where("id = ?", id).Delete(&application).Error
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Success to delete application",
		Data: responses.ApplicationStoreResponse{
			ID:             application.ID,
			UserID:         application.UserID,
			ProjectID:      application.ProjectID,
			CuriculumVitae: application.CuriculumVitae,
			Portofolio:     application.Portofolio,
			Status:         application.Status,
		},
	})
}
