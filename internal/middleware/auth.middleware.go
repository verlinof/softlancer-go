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

func AuthLogin(c *gin.Context) {
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
			c.JSON(http.StatusUnauthorized, errResponse)
			return
		}

		// Check User
		err := database.DB.Where("id = ?", claims["sub"]).First(&user).Error
		if user.ID == "" { // Assuming user.ID is of type string
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		// Set User To Context
		c.Set("user", user.ID)
		// Next
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
	}
}

func AuthAdmin(c *gin.Context) {
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
			c.JSON(http.StatusUnauthorized, errResponse)
			return
		}

		// Check User
		database.DB.Where("id = ?", claims["sub"]).First(&user)
		if user.ID == "" { // Assuming user.ID is of type uint or similar
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
			return
		}

		// Check Role
		if !user.IsAdmin {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
			return
		}

		// Set User To Context
		c.Set("user", user.ID)
		// Next
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
	}
}
