package user_controller

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/verlinof/softlancer-go/database"
	"github.com/verlinof/softlancer-go/models"
	"github.com/verlinof/softlancer-go/requests"
	"github.com/verlinof/softlancer-go/responses"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	userReq := new(requests.UserRequest)

	//Input validation
	if errReq := c.ShouldBind(&userReq); errReq != nil { //ini auto buat bind ataupun postform
		errorResponse := responses.ErrorResponse{
			Status:     "error",
			StatusCode: 400,
			Error:      errReq.Error(),
		}

		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	//Check if the email already exist
	userEmailExist := new(models.User)
	database.DB.Table("users").Where("email = ?", userReq.Email).First(&userEmailExist)

	if userEmailExist.Id != nil {
		errorResponse := responses.ErrorResponse{
			Status:     "error",
			StatusCode: 400,
			Error:      "Email already exist",
		}

		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Encrypt the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		errorResponse := responses.ErrorResponse{
			Status:     "error",
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	hashedPasswordStr := string(hashedPassword)

	//Create User
	user := models.User{
		Name:     &userReq.Name,
		Address:  &userReq.Address,
		Email:    &userReq.Email,
		Password: &hashedPasswordStr,
	}

	err = database.DB.Create(&user).Error

	if err != nil {
		errorResponse := responses.ErrorResponse{
			Status:     "error",
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	userResponse := responses.UserResponse{
		Id:      user.Id,
		Name:    user.Name,
		Address: user.Address,
		Email:   user.Email,
	}
	successResponse := responses.SuccessResponse{
		Status:  "success",
		Message: "Success",
		Data:    userResponse,
	}

	c.JSON(http.StatusOK, successResponse)
}

func Login(c *gin.Context) {
	var userReq requests.LoginRequest
	var user models.User
	var errResponse responses.ErrorResponse
	// Get the email and pass from req body
	if err := (c.ShouldBind(&userReq)); err != nil {
		errResponse = responses.ErrorResponse{
			Status:     "error",
			StatusCode: 400,
			Error:      err.Error(),
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	// Find requested User
	database.DB.Table("users").Where("email = ?", userReq.Email).First(&user)
	if *user.Id == 0 {
		errResponse = responses.ErrorResponse{
			Status:     "error",
			StatusCode: 400,
			Error:      "Invalid Credentials",
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	// Compare the password
	err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(userReq.Password))

	if err != nil {
		errResponse = responses.ErrorResponse{
			Status:     "error",
			StatusCode: 400,
			Error:      "Invalid Credential",
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,                                   //Subject is User ID
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(), //Token will expire in 7 days
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		errResponse = responses.ErrorResponse{
			Status:     "error",
			StatusCode: 500,
			Error:      "Failed to create token",
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, errResponse)
		return
	}

	successResponse := responses.LoginResponse{
		Status:  "success",
		Message: "Success",
		Token:   tokenString,
	}

	c.JSON(http.StatusOK, successResponse)
}

func Show(c *gin.Context) {
	//Get ID
	id := c.Param("id")
	user := new(models.User)
	err := database.DB.First(&user, id)

	//Error handling
	if user.Id == nil {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "User not found",
		})
		return
	}

	if err.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	response := responses.UserResponse{
		Id:      user.Id,
		Name:    user.Name,
		Address: user.Address,
		Email:   user.Email,
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"data":    response,
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
	if errDb != nil {
		c.JSON(500, gin.H{
			"message": errDb.Error(),
		})
		return
	}

	//Update User
	errUpdate := database.DB.Model(&user).Updates(&userReq).Error
	if errUpdate != nil {
		c.JSON(500, gin.H{
			"message": "Failed to update user",
			"error":   errUpdate.Error(),
		})
		return
	}

	//Success
	c.JSON(200, gin.H{
		"message": "Success",
		"data":    user,
	})
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	user := new(models.User)

	//Find User
	database.DB.Table("users").Where("id = ?", id).First(&user)
	errDb := database.DB.Table("users").Where("id = ?", id).Delete(&user).Error
	if errDb != nil {
		c.JSON(500, gin.H{
			"message": errDb.Error(),
		})
		return
	}

	response := responses.UserResponse{
		Id:      user.Id,
		Name:    user.Name,
		Address: user.Address,
		Email:   user.Email,
	}

	//Success
	c.JSON(200, gin.H{
		"message": "Success",
		"data":    response,
	})
}

func IndexPaginate(c *gin.Context) {
	page := c.Query("page")
	if page == "" {
		page = "1"
	}
	perPage := c.Query("per_page")
	if perPage == "" {
		perPage = "10"
	}
	users := new([]models.User) //Buat array

	perPageInt, _ := strconv.Atoi(perPage)
	pageInt, _ := strconv.Atoi(page)
	if pageInt < 1 {
		pageInt = 1
	}

	//Get all users
	err := database.DB.Table("users").Offset((pageInt - 1) * perPageInt).Limit(perPageInt).Find(&users).Error

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "Success",
		"data":     users,
		"page":     pageInt,
		"per_page": perPageInt,
	})
}
