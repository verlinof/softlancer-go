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
	"github.com/verlinof/softlancer-go/pkg"
)

type ProjectController struct {
	emailService   *services.EmailService
	projectService *services.ProjectService
}

func NewProjectController() *ProjectController {
	return &ProjectController{
		emailService:   services.NewEmailService(),
		projectService: services.NewProjectService(),
	}
}

func (e *ProjectController) IndexAdmin(c *gin.Context) {
	var projects []models.ProjectDetail
	projects, err := e.projectService.GetAllProjects(c.Request.Context())
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	for i := range projects {
		logoPath := *pkg.PrefixBaseUrl(projects[i].CompanyLogo)
		projects[i].CompanyLogo = logoPath
	}

	successRes := responses.SuccessResponse{
		Message: "Success",
		Data:    projects,
	}

	c.JSON(http.StatusOK, successRes)
}

func (e *ProjectController) Index(c *gin.Context) {
	var projects []models.ProjectDetail
	projects, err := e.projectService.GetOpenProjects(c.Request.Context())
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	for i := range projects {
		logoPath := *pkg.PrefixBaseUrl(projects[i].CompanyLogo)
		projects[i].CompanyLogo = logoPath
	}

	successRes := responses.SuccessResponse{
		Message: "Success",
		Data:    projects,
	}

	c.JSON(http.StatusOK, successRes)
}

func (e *ProjectController) Show(c *gin.Context) {
	var err error
	var project *models.ProjectDetail

	id := c.Param("id")
	project, err = e.projectService.GetProjectByID(c.Request.Context(), id)
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 404,
			Error:      err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	//Prefix the image url
	project.CompanyLogo = *pkg.PrefixBaseUrl(project.CompanyLogo)

	successRes := responses.SuccessResponse{
		Message: "Success",
		Data:    project, // Mengambil hasil pertama dari slice sebagai response
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
	err = validations.ValidateCreateProject(&projectReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ErrorResponse{
			StatusCode: 400,
			Error:      err.Error(),
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
	err = e.projectService.CreateProject(c.Request.Context(), &project)
	if err != nil {
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
			CompanyID:          project.CompanyID,
			RoleID:             project.RoleID,
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
	// var oldProject models.Project

	id := c.Param("id")

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
	err = validations.ValidateUpdateProject(&projectReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ErrorResponse{
			StatusCode: 400,
			Error:      err,
		})
		return
	}

	// Simpan project ke database
	err = e.projectService.UpdateProject(c.Request.Context(), id, &projectReq)
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	//Cari project berdasarkan ID
	err = database.DB.Table("projects").Where("id = ?", id).First(&project).Error

	if err != nil && strings.Contains(err.Error(), "record not found") {
		errResponse := responses.ErrorResponse{
			StatusCode: 404,
			Error:      "record not found",
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	// Mengembalikan response sukses
	successRes := responses.SuccessResponse{
		Message: "Success",
		Data: responses.ProjectResponse{
			ID:                 project.ID,
			ProjectTitle:       project.ProjectTitle,
			CompanyID:          project.CompanyID,
			RoleID:             project.RoleID,
			ProjectDescription: project.ProjectDescription,
			JobType:            project.JobType,
			Status:             project.Status,
		},
	}

	c.JSON(http.StatusOK, successRes)
}

func (e *ProjectController) Destroy(c *gin.Context) {
	var err error
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
	err = e.projectService.DeleteProject(c.Request.Context(), id)

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
