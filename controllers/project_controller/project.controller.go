package project_controller

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

func Index(c *gin.Context) {
	var projectRes []responses.ProjectResponse
	err := database.DB.Table("projects").
		Select("id, project_title, project_description, job_type, status").
		Scan(&projectRes).Error
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	message := "Success"
	if len(projectRes) == 0 {
		message = "Projects data is empty"
	}

	successRes := responses.SuccessResponse{
		Message: message,
		Data:    projectRes,
	}

	c.JSON(http.StatusOK, successRes)
}

func Show(c *gin.Context) {
	var err error
	var projectRes responses.ProjectResponse

	id := c.Param("id")
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

	successRes := responses.SuccessResponse{
		Message: "Success",
		Data:    projectRes,
	}

	c.JSON(http.StatusOK, successRes)
}

func Store(c *gin.Context) {
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
		ProjectTitle:       &projectReq.ProjectTitle,
		ProjectDescription: &projectReq.ProjectDescription,
		CompanyID:          projectReq.CompanyId,
		RoleID:             projectReq.RoleId,
		JobType:            &projectReq.JobType,
		Status:             &projectReq.Status,
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

func Update(c *gin.Context) {
	var err error
	var projectReq requests.ProjectRequest
	var project models.Project

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
	validationErr := validations.ValidateUpdateProject(&projectReq)
	if len(validationErr) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ErrorResponse{
			StatusCode: 400,
			Error:      validationErr,
		})
		return
	}

	project = models.Project{
		ProjectTitle:       &projectReq.ProjectTitle,
		ProjectDescription: &projectReq.ProjectDescription,
		CompanyID:          projectReq.CompanyId,
		RoleID:             projectReq.RoleId,
		JobType:            &projectReq.JobType,
		Status:             &projectReq.Status,
	}

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

func Delete(c *gin.Context) {
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
