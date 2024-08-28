package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/verlinof/softlancer-go/database"
	"github.com/verlinof/softlancer-go/models"
	"github.com/verlinof/softlancer-go/requests"
	"github.com/verlinof/softlancer-go/responses"
	"github.com/verlinof/softlancer-go/validations"
)

type ProjectController struct{}

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
		Scan(&response).Error

	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	if len(response) == 0 {
		errResponse := responses.ErrorResponse{
			StatusCode: 404,
			Error:      "Project not found",
		}
		c.JSON(http.StatusNotFound, errResponse)
		return
	}

	successRes := responses.SuccessResponse{
		Message: "Success",
		Data:    response[0], // Mengambil hasil pertama dari slice sebagai response
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
	// Mengisi model project berdasarkan projectReq
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
	err = database.DB.Table("projects").Where("id = ?", parsedId).First(&oldProject).Error
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

	project = models.Project{
		ProjectTitle:       projectReq.ProjectTitle,
		ProjectDescription: projectReq.ProjectDescription,
		CompanyID:          projectReq.CompanyId,
		RoleID:             projectReq.RoleId,
		JobType:            projectReq.JobType,
		Status:             projectReq.Status,
	}

	project.ID = uint(parsedId)

	// Simpan project ke database
	if err = database.DB.Where("id = ?", parsedId).Updates(&project).Error; err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

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

	// Mengisi model project berdasarkan projectReq
	parsedId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Invalid ID",
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// Cari project berdasarkan ID
	err = database.DB.Table("projects").
		Select("id, project_title, project_description, job_type, status").
		Where("id = ?", id).
		Scan(&projectRes).Error

	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	if projectRes.ID == 0 {
		errResponse := responses.ErrorResponse{
			StatusCode: 404,
			Error:      "Project not found",
		}
		c.JSON(http.StatusNotFound, errResponse)
		return
	}

	//Delete the project
	err = database.DB.Delete(&project, parsedId).Error

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
