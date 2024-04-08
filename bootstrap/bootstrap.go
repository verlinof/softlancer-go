package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/verlinof/restful-api-golang/config/app_config"
	"github.com/verlinof/restful-api-golang/database"
	"github.com/verlinof/restful-api-golang/routes"
)

//Function capital agar dapat diakses diluar package
func Bootstrap() {
	database.ConnectDatabase()
	app := gin.Default()

	routes.InitRoute(app)

	app.Run(app_config.PORT)
}