package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/verlinof/softlancer-go/database"
	"github.com/verlinof/softlancer-go/models"
	"github.com/verlinof/softlancer-go/responses"
)

func AuthLogin(c *gin.Context) {
	errResponse := responses.ErrorResponse{
		Status:     "error",
		StatusCode: 401,
		Error:      "Unauthorized",
	}

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
		var user models.User
		database.DB.First(&user, claims["sub"])
		if *user.Id == 0 { // Assuming user.ID is of type uint or similar
			log.Fatal(user.Id)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
			return
		}

		// Set User To Context
		c.Set("user", user)
		// Next
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
	}
}
