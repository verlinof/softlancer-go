package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // How long preflight requests can be cached
	}))

	//Init routes
	routes.InitRoute(app)

	//Run server
	err = app.Run(app_config.PORT)
	if err != nil {
		log.Println(err.Error())
	}
}
