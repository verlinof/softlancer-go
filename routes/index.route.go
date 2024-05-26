package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/verlinof/restful-api-golang/config/app_config"
	"github.com/verlinof/restful-api-golang/controllers/file_controller"
	"github.com/verlinof/restful-api-golang/controllers/user_controller"
	"github.com/verlinof/restful-api-golang/middleware"
)

func InitRoute(app *gin.Engine) {
	route := app
	
	// Static Asset
	route.Static(app_config.STATIC_PATH, app_config.STATIC_DIR)

	//Route User
	userRoute := route.Group("/users") //Untuk grouping Route atau bisa disebut Prefix
	userRoute.GET("/", user_controller.Index)
	userRoute.GET("/paginate", user_controller.IndexPaginate)
	userRoute.GET("/:id", user_controller.Show)
	userRoute.POST("", user_controller.Store)
	userRoute.PUT("/:id", user_controller.Update)
	userRoute.DELETE("/:id", user_controller.Delete)

	//Route File
	fileRoute := route.Group("/file", middleware.AuthMiddleware) //Kalau nambahin Middleware
	fileRoute.POST("/", file_controller.HandleUploadFile)
	fileRoute.DELETE("/:filename", file_controller.HandleRemoveFile)
}