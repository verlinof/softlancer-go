package bootstrap

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/verlinof/restful-api-golang/config/app_config"
	"github.com/verlinof/restful-api-golang/config/db_config"
	"github.com/verlinof/restful-api-golang/database"
	"github.com/verlinof/restful-api-golang/routes"
)

//Function capital agar dapat diakses diluar package
func Bootstrap() {
	//Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	//Load config
	app_config.InitAppConfig()
	db_config.InitDatabaseConfig()

	//Load database
	database.ConnectDatabase()
	//Init GIN ENGINE
	app := gin.Default()

	routes.InitRoute(app)

	app.Run(app_config.PORT)
}