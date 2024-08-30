package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/verlinof/softlancer-go/config/app_config"
	"github.com/verlinof/softlancer-go/config/db_config"
	"github.com/verlinof/softlancer-go/database"
	"github.com/verlinof/softlancer-go/routes"
)

func main() {
	//Load .env file
	err := godotenv.Load("./.env")
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
