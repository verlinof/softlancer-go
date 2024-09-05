package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/models"
	"github.com/verlinof/softlancer-go/internal/requests"
	"github.com/verlinof/softlancer-go/internal/responses"
)

type ReferenceController struct{}

func (e *ReferenceController) Index(c *gin.Context) {
	var response []responses.ReferenceResponse

	err := database.DB.Table("references").Find(&response).Error
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Internal Server Error",
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Success",
		Data:    response,
	})
}

func (e *ReferenceController) Show(c *gin.Context) {
	var response responses.ReferenceResponse
	var err error

	id := c.Param("id")

	err = database.DB.Table("references").Where("id = ?", id).First(&response).Error
	if err != nil && strings.Contains(err.Error(), "record not found") {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Reference not found",
		}
		c.JSON(http.StatusNotFound, errResponse)
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Success",
		Data:    response,
	})
}

func (e *ReferenceController) Store(c *gin.Context) {
	var request requests.CreateReferenceRequest
	var err error

	userId, _ := c.Get("user")

	err = c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			StatusCode: 400,
			Error:      "Invalid request body",
		})
		return
	}

	reference := models.Reference{
		UserID: userId.(string),
		RoleID: request.RoleID,
	}

	err = database.DB.Table("references").Create(&reference).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Success",
		Data: responses.ReferenceResponse{
			ID:     reference.ID,
			UserID: reference.UserID,
			RoleID: reference.RoleID,
		},
	})
}

func (e *ReferenceController) Destroy(c *gin.Context) {
	var reference models.Reference
	var err error

	id := c.Param("id")

	err = database.DB.Table("references").Where("id = ?", id).First(&reference).Error
	if err != nil && strings.Contains(err.Error(), "record not found") {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			StatusCode: 404,
			Error:      "Reference not found",
		})
		return
	}

	err = database.DB.Table("references").Where("id = ?", id).Delete(&reference).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Success to delete reference",
		Data: responses.ReferenceResponse{
			ID:     reference.ID,
			UserID: reference.UserID,
			RoleID: reference.RoleID,
		},
	})
}
