package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	route := app
	route.GET("/", func(c *gin.Context) {
		isValidated := false

		if(!isValidated) {
			c.JSON(200, gin.H{
				"message": "Bad Request, some field not valid",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
		return
	})

	route.Run(":8000")
}