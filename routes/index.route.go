package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/verlinof/softlancer-go/config/app_config"
	"github.com/verlinof/softlancer-go/controllers/file_controller"
	"github.com/verlinof/softlancer-go/controllers/user_controller"
	"github.com/verlinof/softlancer-go/middleware"
)

func InitRoute(app *gin.Engine) {
	// Static Asset
	app.Static(app_config.STATIC_PATH, app_config.STATIC_DIR)

	// Base route group with /api prefix
	api := app.Group("/api")

	//Auth Routes
	authRoute := api.Group("/auth")
	authRoute.POST("/login", user_controller.Login)
	authRoute.POST("/register", user_controller.Register)

	// User routes
	userRoute := api.Group("/users") // Grouping routes with /users prefix
	userRoute.GET("/register", user_controller.Register)

	api.GET("/tes", middleware.AuthLogin, user_controller.Tes)

	// File routes
	fileRoute := api.Group("/file", middleware.AuthLogin) // Grouping routes with /file prefix and middleware
	fileRoute.POST("/", file_controller.HandleUploadFile)
	fileRoute.DELETE("/:filename", file_controller.HandleRemoveFile)
}
