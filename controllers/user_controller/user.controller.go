package user_controller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/verlinof/softlancer-go/database"
	"github.com/verlinof/softlancer-go/models"
	"github.com/verlinof/softlancer-go/requests"
	"github.com/verlinof/softlancer-go/responses"
	"golang.org/x/crypto/bcrypt"
)

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
	if *user.ID == 0 {
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

	if userEmailExist.ID != nil {
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
		ID:      user.ID,
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

func Profile(c *gin.Context) {
	userId, exists := c.Get("user")
	if !exists {
		errorResponse := responses.ErrorResponse{
			Status:     "error",
			StatusCode: 500,
			Error:      "Internal Server Error",
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	var user models.User
	if err := database.DB.First(&user, userId); err.Error != nil {
		errorResponse := responses.ErrorResponse{
			Status:     "error",
			StatusCode: 500,
			Error:      "User not found",
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	userResponse := responses.UserResponse{
		ID:      user.ID,
		Name:    user.Name,
		Address: user.Address,
		Email:   user.Email,
	}
	successReponse := responses.SuccessResponse{
		Status:  "success",
		Message: "Success to get user profile",
		Data:    userResponse,
	}

	c.JSON(http.StatusOK, successReponse)
}

func Update(c *gin.Context) {
	id, _ := c.Get("user")
	// Initialize Validator
	// var validate *validator.Validate
	user := new(models.User)
	userReq := new(requests.UpdateUserRequest)
	userRes := new(responses.UserResponse)

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
		errorResponse := responses.ErrorResponse{
			Status:     "error",
			StatusCode: 500,
			Error:      "Internal Server Error",
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	//Update User
	errUpdate := database.DB.Table("users").Where("id = ?", id).Updates(&userReq).Error
	if errUpdate != nil {
		errorResponse := responses.ErrorResponse{
			Status:     "error",
			StatusCode: 500,
			Error:      "Error Updating User",
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	userRes.ID = user.ID
	userRes.Name = user.Name
	userRes.Address = user.Address
	userRes.Email = user.Email

	//Success
	c.JSON(200, gin.H{
		"message": "Success",
		"data":    userRes,
	})
}
