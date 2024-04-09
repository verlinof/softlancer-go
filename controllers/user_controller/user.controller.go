package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/verlinof/restful-api-golang/database"
	"github.com/verlinof/restful-api-golang/models"
)

func Index(c *gin.Context) {
	users := new([]models.User) //Buat array

	//Get all users
	// database.DB.Find(&users)
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
	
	err := database.DB.First(&models.User{}, id).Error

	if(err != nil) {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"data" : models.User{},
	})
}