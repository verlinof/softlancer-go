package project_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/verlinof/softlancer-go/database"
	"github.com/verlinof/softlancer-go/responses"
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
