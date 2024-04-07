package bootstrap

import (
	"github.com/gin-gonic/gin"
	app_config "github.com/verlinof/restful-api-golang/config"
	"github.com/verlinof/restful-api-golang/routes"
)

//Function capital agar dapat diakses diluar package
func Bootstrap() {
	
	app := gin.Default()

	routes.InitRoute(app)

	app.Run(app_config.PORT)
}