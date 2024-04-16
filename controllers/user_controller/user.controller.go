package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/verlinof/restful-api-golang/database"
	"github.com/verlinof/restful-api-golang/models"
	"github.com/verlinof/restful-api-golang/requests"
	"github.com/verlinof/restful-api-golang/responses"
)

func Index(c *gin.Context) {
	users := new([]models.User) //Buat array

	//Get all users
	// Cara buat spesifik table
	err := database.DB.Table("users").Find(&users)

	if(err.Error != nil) {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"data" : users,
	})
}

func Show(c *gin.Context) {
	//Get ID
	id := c.Param("id")
	user := new(models.User)
	err := database.DB.First(&user, id)

	//Error handling
	if(user.Id == nil) {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "User not found",
		})
		return
	}

	if(err.Error != nil) {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	response := responses.UserResponse{
		Id: user.Id,
		Name: user.Name,
		Address: user.Address,
		Email: user.Email,
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"data" : response,
	})
}

func Store(c *gin.Context) {
	userReq := new(requests.UserRequest)
	if errReq := c.ShouldBind(&userReq); errReq != nil { //ini auto buat bind ataupun postform
		c.JSON(400, gin.H{
			"message": errReq.Error(),
		})
		return
	} 

	user := models.User{
		Name: &userReq.Name,
		Address: &userReq.Address,
		Email: &userReq.Email,
		Password: &userReq.Password,
		Born_date: &userReq.Born_date,
	}

	err := database.DB.Create(&user).Error

	if(err != nil) {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"data" : user,
	})
}