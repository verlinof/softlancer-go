package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/models"
	"github.com/verlinof/softlancer-go/internal/requests"
	"github.com/verlinof/softlancer-go/internal/responses"
	"github.com/verlinof/softlancer-go/internal/services"
	"github.com/verlinof/softlancer-go/internal/validations"
)

type ProjectController struct {
	emailService *services.EmailService
}

func (e *ProjectController) Init() {
	if e.emailService == nil {
		e.emailService = services.NewEmailService()
	}
}

func (e *ProjectController) IndexAdmin(c *gin.Context) {
	var response []map[string]interface{}
	err := database.DB.Table("projects").
		Select(`
			projects.id, 
			projects.project_title, 
			projects.project_description, 
			projects.job_type, 
			projects.status,
			roles.role_name,
			companies.company_name, 
			companies.company_description, 
			companies.company_logo
		`).
		Joins("JOIN companies ON projects.company_id = companies.id").
		Joins("JOIN roles ON projects.role_id = roles.id").
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
		message = "Projects data is empty"
	}

	successRes := responses.SuccessResponse{
		Message: message,
		Data:    response,
	}

	c.JSON(http.StatusOK, successRes)
}

func (e *ProjectController) Index(c *gin.Context) {
	var response []map[string]interface{}
	err := database.DB.Table("projects").
		Select(`
		projects.id, 
		projects.project_title, 
		projects.project_description, 
		projects.job_type, 
		projects.status,
		roles.role_name,
		companies.company_name, 
		companies.company_description, 
		companies.company_logo
	`).
		Joins("JOIN companies ON projects.company_id = companies.id").
		Joins("JOIN roles ON projects.role_id = roles.id").
		Where("projects.status = ?", "open").
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
		message = "Projects data is empty"
	}

	successRes := responses.SuccessResponse{
		Message: message,
		Data:    response,
	}

	c.JSON(http.StatusOK, successRes)
}

func (e *ProjectController) Show(c *gin.Context) {
	var err error
	var response []map[string]interface{}

	id := c.Param("id")
	err = database.DB.Table("projects").
		Select(`
			projects.id, 
			projects.project_title, 
			projects.project_description, 
			projects.job_type, 
			projects.status,
			roles.role_name,
			companies.company_name, 
			companies.company_description, 
			companies.company_logo
	`).
		Joins("JOIN companies ON projects.company_id = companies.id").
		Joins("JOIN roles ON projects.role_id = roles.id").
		Where("projects.id = ?", id).
		First(&response).Error

	if err != nil && strings.Contains(err.Error(), "record not found") {
		errResponse := responses.ErrorResponse{
			StatusCode: 404,
			Error:      "Project not found",
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	successRes := responses.SuccessResponse{
		Message: "Success",
		Data:    response, // Mengambil hasil pertama dari slice sebagai response
	}

	c.JSON(http.StatusOK, successRes)
}

func (e *ProjectController) Store(c *gin.Context) {
	var err error
	var projectReq requests.ProjectRequest
	var project models.Project

	// Bind request body ke struct ProjectRequest
	if err = c.ShouldBind(&projectReq); err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 400,
			Error:      "Invalid request body",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	// Validasi input dari user
	validationErr := validations.ValidateCreateProject(&projectReq)
	if len(validationErr) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ErrorResponse{
			StatusCode: 400,
			Error:      validationErr,
		})
		return
	}

	// Mengisi model project berdasarkan projectReq
	project = models.Project{
		ProjectTitle:       projectReq.ProjectTitle,
		ProjectDescription: projectReq.ProjectDescription,
		CompanyID:          projectReq.CompanyId,
		RoleID:             projectReq.RoleId,
		JobType:            projectReq.JobType,
		Status:             projectReq.Status,
	}

	// Simpan project ke database
	if err = database.DB.Create(&project).Error; err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// Menjalankan perintah sendEmailJob
	go e.emailService.SendEmail(project.RoleID, project.ProjectTitle)

	// Mengembalikan response sukses
	successRes := responses.SuccessResponse{
		Message: "Success",
		Data: responses.ProjectResponse{
			ID:                 project.ID,
			ProjectTitle:       project.ProjectTitle,
			ProjectDescription: project.ProjectDescription,
			JobType:            project.JobType,
			Status:             project.Status,
		},
	}

	c.JSON(http.StatusOK, successRes)
}

func (e *ProjectController) Update(c *gin.Context) {
	var err error
	var projectReq requests.ProjectRequest
	var project models.Project
	var oldProject models.Project

	id := c.Param("id")

	//Find the old data
	err = database.DB.Table("projects").Where("id = ?", id).First(&oldProject).Error
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 404,
			Error:      "Project not found",
		}
		c.JSON(http.StatusNotFound, errResponse)
		return
	}

	// Bind request body ke struct ProjectRequest
	if err = c.ShouldBind(&projectReq); err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 400,
			Error:      "Invalid request body",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	// Validasi input dari user
	validationErr := validations.ValidateUpdateProject(&projectReq)
	if len(validationErr) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ErrorResponse{
			StatusCode: 400,
			Error:      validationErr,
		})
		return
	}

	// Simpan project ke database
	if err = database.DB.Table("projects").Where("id = ?", id).Updates(&projectReq).Error; err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// Mencari data yang telah diupdate
	database.DB.Table("projects").Where("id = ?", id).First(&project)

	// Mengembalikan response sukses
	successRes := responses.SuccessResponse{
		Message: "Success",
		Data: responses.ProjectResponse{
			ID:                 project.ID,
			ProjectTitle:       project.ProjectTitle,
			ProjectDescription: project.ProjectDescription,
			JobType:            project.JobType,
			Status:             project.Status,
		},
	}

	c.JSON(http.StatusOK, successRes)
}

func (e *ProjectController) Destroy(c *gin.Context) {
	var err error
	var project models.Project
	var projectRes responses.ProjectResponse

	id := c.Param("id")

	// Cari project berdasarkan ID
	err = database.DB.Table("projects").
		Where("id = ?", id).
		First(&projectRes).Error

	if err != nil && strings.Contains(err.Error(), "record not found") {
		errResponse := responses.ErrorResponse{
			StatusCode: 404,
			Error:      "Project not found",
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	//Delete the project
	err = database.DB.Where("id = ?", id).Delete(&project).Error

	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// Mengembalikan response sukses
	successRes := responses.SuccessResponse{
		Message: "Success Deleting Project",
		Data:    projectRes,
	}

	c.JSON(http.StatusOK, successRes)
}
