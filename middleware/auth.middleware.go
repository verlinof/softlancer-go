package middleware

import "github.com/gin-gonic/gin"

func AuthMiddleware(c *gin.Context) {
	c.GetHeader("Authorization")
	c.Next()
}