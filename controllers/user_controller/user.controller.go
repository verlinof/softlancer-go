package user_controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/verlinof/softlancer-go/database"
	"github.com/verlinof/softlancer-go/models"
	"github.com/verlinof/softlancer-go/requests"
	"github.com/verlinof/softlancer-go/responses"
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

	userEmailExist := new(models.User)
	database.DB.Table("users").Where("email = ?", userReq.Email).First(&userEmailExist)

	if(userEmailExist.Id != nil) {
		c.JSON(400, gin.H{
			"message": "Email already exist",
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
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"data" : user,
	})
}

func Update(c *gin.Context) {
	id := c.Param("id")
	user := new(models.User)
	userReq := new(requests.UserRequest)

	//If error with the Request Body
	if errReq := c.ShouldBind(&userReq); errReq != nil {
		c.JSON(400, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	//Find User
	errDb := database.DB.Table("users").Where("id = ?", id).First(&user).Error
	if(errDb != nil) {
		c.JSON(500, gin.H{
			"message": errDb.Error(),
		})
		return
	}

	//Update User
	errUpdate := database.DB.Model(&user).Updates(&userReq).Error
	if(errUpdate != nil) {
		c.JSON(500, gin.H{
			"message": "Failed to update user",
			"error" : errUpdate.Error(),
		})
		return
	}

	//Success
	c.JSON(200, gin.H{
		"message": "Success",
		"data" : user,
	})
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	user := new(models.User)

	//Find User
	database.DB.Table("users").Where("id = ?", id).First(&user)
	errDb := database.DB.Table("users").Where("id = ?", id).Delete(&user).Error
	if(errDb != nil) {
		c.JSON(500, gin.H{
			"message": errDb.Error(),
		})
		return
	}

	response :=  responses.UserResponse{
		Id: user.Id,
		Name: user.Name,
		Address: user.Address,
		Email: user.Email,
	}

	//Success
	c.JSON(200, gin.H{
		"message": "Success",
		"data" : response,
	})
}

func IndexPaginate(c *gin.Context) {
	page := c.Query("page")
	if(page == "" ){
		page = "1"
	}
	perPage := c.Query("per_page")
	if(perPage == "" ){
		perPage = "10"
	}
	users := new([]models.User) //Buat array

	perPageInt, _ := strconv.Atoi(perPage)
	pageInt, _ := strconv.Atoi(page)
	if(pageInt < 1) {
		pageInt = 1
	}

	//Get all users
	err := database.DB.Table("users").Offset((pageInt - 1) * perPageInt).Limit(perPageInt).Find(&users).Error

	if(err != nil) {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"data" : users,
		"page": pageInt,
		"per_page": perPageInt,
	})
}