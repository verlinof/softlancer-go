package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/models"
	"github.com/verlinof/softlancer-go/internal/requests"
	"github.com/verlinof/softlancer-go/internal/responses"
	"github.com/verlinof/softlancer-go/internal/services"
	"github.com/verlinof/softlancer-go/internal/validations"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	Service *services.UserService
}

func (e *UserController) Init() {
	if e.Service == nil {
		e.Service = services.NewUserService()
	}
}

func (e *UserController) Index(c *gin.Context) {
	var userRes []responses.UserResponse
	userRes, err := e.Service.GetUsers(c.Request.Context())
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Success",
		Data:    userRes,
	})
}

func (e *UserController) Login(c *gin.Context) {
	var userReq requests.LoginRequest
	var user *models.User
	var errResponse responses.ErrorResponse

	// Get the email and pass from req body
	if err := c.ShouldBind(&userReq); err != nil {
		errResponse = responses.ErrorResponse{
			StatusCode: 400,
			Error:      "Invalid Request Body",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	// Validate user input
	validationErr := validations.ValidateLogin(&userReq)
	if len(validationErr) > 0 {
		errResponse = responses.ErrorResponse{
			StatusCode: 400,
			Error:      validationErr,
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	// Find requested User
	database.DB.Table("users").Where("email = ?", userReq.Email).First(&user)
	if user.ID == "" {
		errResponse = responses.ErrorResponse{
			StatusCode: 400,
			Error:      "Invalid Credentials",
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	// Compare the password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password))
	if err != nil {
		errResponse = responses.ErrorResponse{
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
			StatusCode: 500,
			Error:      "Failed to create token",
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, errResponse)
		return
	}

	c.JSON(http.StatusOK, responses.LoginResponse{
		Message: "Success",
		Token:   tokenString,
	})
}

func (e *UserController) Register(c *gin.Context) {
	var err error
	userReq := new(requests.UserRequest)

	//Input validation
	if err := c.ShouldBind(&userReq); err != nil { //ini auto buat bind ataupun postform
		errorResponse := responses.ErrorResponse{
			StatusCode: 400,
			Error:      "Invalid request body",
		}

		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	validationErr := validations.ValidateRegister(userReq)
	if len(validationErr) > 0 {
		errorResponse := responses.ErrorResponse{
			StatusCode: 400,
			Error:      validationErr,
		}

		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Encrypt the password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	hashedPasswordStr := string(hashedPassword)

	//Create User
	user := models.User{
		Name:     userReq.Name,
		Address:  userReq.Address,
		Email:    userReq.Email,
		Password: hashedPasswordStr,
	}

	err = database.DB.Create(&user).Error
	if err != nil {
		errorResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	//Create User Response
	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Success",
		Data: responses.UserResponse{
			ID:      user.ID,
			Name:    user.Name,
			Address: user.Address,
			Email:   user.Email,
		},
	})
}

func (e *UserController) Profile(c *gin.Context) {
	//Get User id from Middleware
	userId, exists := c.Get("user")
	var user models.User

	if !exists {
		errorResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Internal Server Error",
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		errorResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Internal Server Error",
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Success to get user profile",
		Data: responses.UserResponse{
			ID:      user.ID,
			Name:    user.Name,
			Address: user.Address,
			Email:   user.Email,
		},
	})
}
