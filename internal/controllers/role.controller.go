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
)

type RoleController struct{}

func (e *RoleController) Index(c *gin.Context) {
	var response []models.Role

	err := database.DB.Table("roles").Find(&response).Error

	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	successRes := responses.SuccessResponse{
		Message: "Success",
		Data:    response,
	}

	c.JSON(http.StatusOK, successRes)
}

func (e *RoleController) Show(c *gin.Context) {
	var response models.Role

	id := c.Param("id")

	err := database.DB.Table("roles").Where("id = ?", id).First(&response).Error
	if err != nil && strings.Contains(err.Error(), "record not found") {
		errResponse := responses.ErrorResponse{
			StatusCode: 404,
			Error:      "Role not found",
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	successRes := responses.SuccessResponse{
		Message: "Success",
		Data:    response,
	}

	c.JSON(http.StatusOK, successRes)
}

func (e *RoleController) Store(c *gin.Context) {
	var request requests.RoleRequest
	var err error

	if err = c.ShouldBind(&request); err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 400,
			Error:      "Invalid request body",
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	validationErr := validations.ValidateCreateRole(&request)
	if len(validationErr) > 0 {
		errResponse := responses.ErrorResponse{
			StatusCode: 400,
			Error:      validationErr,
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	role := models.Role{
		RoleName: request.RoleName,
	}

	err = database.DB.Table("roles").Create(&role).Error
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Internal Server Error",
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	successRes := responses.SuccessResponse{
		Message: "Success to create role",
		Data:    role,
	}
	c.JSON(http.StatusCreated, successRes)
}

func (e *RoleController) Update(c *gin.Context) {
	var request requests.RoleRequest
	var err error
	var role models.Role

	id := c.Param("id")

	if err = c.ShouldBind(&request); err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 400,
			Error:      "Invalid request body",
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	err = database.DB.Table("roles").Where("id = ?", id).First(&role).Error
	if err != nil && strings.Contains(err.Error(), "record not found") {
		errResponse := responses.ErrorResponse{
			StatusCode: 404,
			Error:      "Role not found",
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	validationErr := validations.ValidateCreateRole(&request)
	if len(validationErr) > 0 {
		errResponse := responses.ErrorResponse{
			StatusCode: 400,
			Error:      validationErr,
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	//Update the data
	role.RoleName = request.RoleName
	err = database.DB.Table("roles").Updates(&role).Error
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Internal Server Error",
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	successRes := responses.SuccessResponse{
		Message: "Success to update role",
		Data:    role,
	}
	c.JSON(http.StatusOK, successRes)
}

func (e *RoleController) Destroy(c *gin.Context) {
	var role models.Role

	id := c.Param("id")

	err := database.DB.Table("roles").Where("id = ?", id).First(&role).Error
	if err != nil && strings.Contains(err.Error(), "record not found") {
		errResponse := responses.ErrorResponse{
			StatusCode: 404,
			Error:      "Role not found",
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	err = database.DB.Table("roles").Delete(&role).Error
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Internal Server Error",
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	successRes := responses.SuccessResponse{
		Message: "Success to delete role",
		Data:    role,
	}
	c.JSON(http.StatusOK, successRes)
}
