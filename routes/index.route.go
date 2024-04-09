package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/verlinof/restful-api-golang/controllers/user_controller"
)

func InitRoute(app *gin.Engine) {
	route := app

	route.GET("/users", user_controller.Index)
}