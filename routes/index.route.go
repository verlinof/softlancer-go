package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/verlinof/softlancer-go/config/app_config"
	"github.com/verlinof/softlancer-go/controllers"
	"github.com/verlinof/softlancer-go/middleware"
)

func InitRoute(app *gin.Engine) {
	//Controller
	var userController *controllers.UserController
	var projectController *controllers.ProjectController
	var companyController *controllers.CompanyController

	// Static Asset
	app.Static(app_config.STATIC_PATH, app_config.STATIC_DIR)

	// Base route group with /api prefix
	api := app.Group("/api")

	//Auth Routes
	authRoute := api.Group("/auth")
	authRoute.POST("/login", userController.Login)
	authRoute.POST("/register", userController.Register)
	authRoute.GET("/profile", middleware.AuthLogin, userController.Profile)
	// authRoute.PATCH("/update-profile", middleware.AuthLogin, userController.Update)

	// User routes
	userRoute := api.Group("/users") // Grouping routes with /users prefix
	userRoute.GET("/", middleware.AuthAdmin, userController.Index)

	//Projects Route
	projectRoute := api.Group("/projects")
	projectRoute.GET("/", projectController.Index)
	projectRoute.GET("/:id", projectController.Show)
	projectRoute.POST("/", middleware.AuthAdmin, projectController.Store)
	projectRoute.PATCH("/:id", middleware.AuthAdmin, projectController.Update)
	projectRoute.DELETE("/:id", middleware.AuthAdmin, projectController.Destroy)

	//Company Route
	companyRoute := api.Group("/companies")
	companyRoute.GET("/", companyController.Index)
	companyRoute.GET("/:id", companyController.Show)
	companyRoute.POST("/", middleware.AuthAdmin, companyController.Store)
	companyRoute.PATCH("/:id", middleware.AuthAdmin, companyController.Update)
	companyRoute.DELETE("/:id", middleware.AuthAdmin, companyController.Destroy)

	// File routes
	// fileRoute := api.Group("/file", middleware.AuthLogin) // Grouping routes with /file prefix and middleware
	// fileRoute.POST("/", file_controller.HandleUploadFile)
	// fileRoute.DELETE("/:filename", file_controller.HandleRemoveFile)
}
