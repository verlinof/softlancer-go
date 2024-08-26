package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/verlinof/softlancer-go/config/app_config"
	"github.com/verlinof/softlancer-go/controllers"
	"github.com/verlinof/softlancer-go/controllers/file_controller"
	"github.com/verlinof/softlancer-go/middleware"
)

func InitRoute(app *gin.Engine) {
	//Controller
	var projectController *controllers.ProjectController
	var userController *controllers.UserController

	// Static Asset
	app.Static(app_config.STATIC_PATH, app_config.STATIC_DIR)

	// Base route group with /api prefix
	api := app.Group("/api")

	//Auth Routes
	authRoute := api.Group("/auth")
	authRoute.POST("/login", userController.Login)
	authRoute.POST("/register", userController.Register)
	authRoute.GET("/profile", middleware.AuthLogin, userController.Profile)
	// authRoute.PATCH("/update-profile", middleware.AuthLogin, user_controller.Update)

	//Projects Route
	projectRoute := api.Group("/projects")
	projectRoute.GET("/", projectController.Index)
	projectRoute.GET("/:id", projectController.Show)
	projectRoute.POST("/", projectController.Store)
	projectRoute.PATCH("/:id", projectController.Update)
	projectRoute.DELETE("/:id", projectController.Delete)

	// User routes
	userRoute := api.Group("/users") // Grouping routes with /users prefix
	userRoute.GET("/", middleware.AuthAdmin, userController.Index)

	// File routes
	fileRoute := api.Group("/file", middleware.AuthLogin) // Grouping routes with /file prefix and middleware
	fileRoute.POST("/", file_controller.HandleUploadFile)
	fileRoute.DELETE("/:filename", file_controller.HandleRemoveFile)
}
