package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/models"
	"github.com/verlinof/softlancer-go/internal/responses"
)

func ApplicationOwner(c *gin.Context) {
	errResponse := responses.ErrorResponse{
		StatusCode: 401,
		Error:      "Unauthorized",
	}
	var user models.User

	// Get Token From Header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
		return
	}

	applicationId := c.Param("id")

	//Find the application
	var application models.Application
	err := database.DB.Table("applications").Where("id = ?", applicationId).First(&application).Error
	if err != nil {
		errResponse := responses.ErrorResponse{
			StatusCode: 404,
			Error:      "Application not found",
		}
		c.AbortWithStatusJSON(http.StatusNotFound, errResponse)
		return
	}

	// Remove "Bearer " prefix from the token string
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Check Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
		return
	}

	// Validate Token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check Expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
			return
		}

		// Check User
		database.DB.Where("id = ?", claims["sub"]).First(&user)
		if user.ID == "" { // Assuming user.ID is of type uint or similar
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
			return
		}

		//Check is it valid user or no
		if application.UserID != user.ID {
			errResponse := responses.ErrorResponse{
				StatusCode: 401,
				Error:      "Unauthorized",
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
			return
		}

		// Set User To Context
		c.Set("user", user.ID)
		c.Set("application", applicationId)
		// Next
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
	}
}
