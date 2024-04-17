package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/verlinof/restful-api-golang/controllers/user_controller"
)

func InitRoute(app *gin.Engine) {
	route := app

	//Route User
	route.GET("/users", user_controller.Index)
	route.GET("/users/paginate", user_controller.IndexPaginate)
	route.GET("/users/:id", user_controller.Show)
	route.POST("/users", user_controller.Store)
	route.PUT("/users/:id", user_controller.Update)
	route.DELETE("/users/:id", user_controller.Delete)
}