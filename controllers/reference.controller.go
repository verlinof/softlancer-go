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

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Invalid ID",
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	err = database.DB.Table("references").First(&response, id).Error
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

	userId, _ := c.Get("user")

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			StatusCode: 400,
			Error:      "Invalid request body",
		})
		return
	}

	reference := models.Reference{
		UserID: userId.(*uint64),
		RoleID: *request.RoleID,
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
			UserID: *reference.UserID,
			RoleID: reference.RoleID,
		},
	})
}

func (e *ReferenceController) Destroy(c *gin.Context) {
	var reference models.Reference

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			StatusCode: 400,
			Error:      "Invalid ID",
		})
		return
	}

	err = database.DB.Table("references").First(&reference, id).Error
	if err != nil && strings.Contains(err.Error(), "record not found") {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			StatusCode: 404,
			Error:      "Reference not found",
		})
		return
	}

	err = database.DB.Table("references").Where("id = ?", id).Delete(reference).Error
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
			UserID: *reference.UserID,
			RoleID: reference.RoleID,
		},
	})
}
