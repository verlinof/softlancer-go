package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/verlinof/softlancer-go/config/app_config"
	"github.com/verlinof/softlancer-go/config/db_config"
	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/routes"
)

func main() {
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

	//Init routes
	routes.InitRoute(app)

	//Run server
	err = app.Run(app_config.PORT)
	if err != nil {
		log.Println(err.Error())
	}
}
