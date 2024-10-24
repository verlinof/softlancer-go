package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/verlinof/softlancer-go/internal/models"
	"github.com/verlinof/softlancer-go/internal/requests"
	"github.com/verlinof/softlancer-go/internal/responses"
	"github.com/verlinof/softlancer-go/internal/services"
	"github.com/verlinof/softlancer-go/internal/validations"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		UserService: services.NewUserService(),
	}
}

// Index return all user data
//
// @Summary Get all user data
// @ID get-all-user
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.SuccessResponse{data=[]responses.UserResponse}
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /users [get]
func (e *UserController) Index(c *gin.Context) {
	var userRes []responses.UserResponse
	users, err := e.UserService.GetUsers(c.Request.Context())
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	//Insert data to userRes
	for _, user := range *users {
		userRes = append(userRes, responses.UserResponse{
			ID:      user.ID,
			Name:    user.Name,
			Address: user.Address,
			Email:   user.Email,
		})
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Success",
		Data:    userRes,
	})
}

// Login user using email and password
// It will return JWT token if the input is valid
// The token will be expired in 7 days
// The response will be in this format :
//
//	{
//	  "message": "Success",
//	  "token": "your-jwt-token"
//	}
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
	if err := validations.ValidateLogin(&userReq); err != nil {
		errResponse = responses.ErrorResponse{
			StatusCode: 400,
			Error:      err.Error(),
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	// Find requested User
	user, err := e.UserService.GetUserbyEmail(c.Request.Context(), userReq.Email)
	if user.ID == "" {
		errResponse = responses.ErrorResponse{
			StatusCode: 400,
			Error:      "Invalid Credentials",
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	if err != nil {
		errResponse = responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Internal Server Error",
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, errResponse)
		return
	}

	// Compare the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password))
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

// Register new user
//
// @Summary Register new user
// @ID register-new-user
// @Tags User
// @Accept  json
// @Produce  json
// @Param request body requests.UserRequest true "Register new user request body"
// @Success 200 {object} responses.SuccessResponse{data=responses.UserResponse}
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /register [post]
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
	user := &models.User{
		Name:     userReq.Name,
		Address:  userReq.Address,
		Email:    userReq.Email,
		Password: hashedPasswordStr,
	}

	user, err = e.UserService.CreateUser(c.Request.Context(), user)
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
	var user *models.User
	if !exists {
		errorResponse := responses.ErrorResponse{
			StatusCode: 500,
			Error:      "Internal Server Error",
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	user, err := e.UserService.GetUserbyID(c.Request.Context(), userId.(string))
	if err != nil {
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
