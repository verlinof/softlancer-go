package user_controller

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	isValidated := false

		if !isValidated {
			c.JSON(200, gin.H{
				"message": "Bad Request, some field not valid",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
		return
}