package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/verlinof/softlancer-go/config/app_config"
	"github.com/verlinof/softlancer-go/internal/controllers"
	"github.com/verlinof/softlancer-go/internal/middleware"
)

func InitRoute(app *gin.Engine) {
	//Controller
	userController := &controllers.UserController{}
	userController.Init()

	projectController := &controllers.ProjectController{}
	projectController.Init()

	var companyController *controllers.CompanyController
	var applicationController *controllers.ApplicationController
	var roleController *controllers.RoleController
	var referenceController *controllers.ReferenceController

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
	projectRoute.GET("/all", middleware.AuthAdmin, projectController.IndexAdmin)
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

	//Application Route
	applicationRoute := api.Group("/applications")
	applicationRoute.GET("/", applicationController.Index)
	applicationRoute.GET("/:id", applicationController.Show)
	applicationRoute.POST("/", middleware.AuthLogin, applicationController.Store)
	applicationRoute.PATCH("/status/:id", middleware.AuthAdmin, applicationController.UpdateStatus)
	applicationRoute.PATCH("/:id", middleware.ApplicationOwner, applicationController.Update)
	applicationRoute.DELETE("/:id", middleware.ApplicationOwner, applicationController.Destroy)

	//Role Route
	roleRoute := api.Group("/roles")
	roleRoute.GET("/", roleController.Index)
	roleRoute.GET("/:id", roleController.Show)
	roleRoute.POST("/", middleware.AuthAdmin, roleController.Store)
	roleRoute.PATCH("/:id", middleware.AuthAdmin, roleController.Update)
	roleRoute.DELETE("/:id", middleware.AuthAdmin, roleController.Destroy)

	//Reference Route
	referenceRoute := api.Group("/references")
	referenceRoute.GET("/", referenceController.Index)
	referenceRoute.GET("/:id", referenceController.Show)
	referenceRoute.POST("/", middleware.AuthLogin, referenceController.Store)
	referenceRoute.DELETE("/:id", middleware.AuthLogin, referenceController.Destroy)
}
