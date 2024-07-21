package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/verlinof/softlancer-go/config/app_config"
	"github.com/verlinof/softlancer-go/config/db_config"
	"github.com/verlinof/softlancer-go/database"
	"github.com/verlinof/softlancer-go/models"
)

func init() {
	err :=  godotenv.Load()
	if(err != nil) {
		log.Println(err)
	}

	//init config
	app_config.InitAppConfig()
	db_config.InitDatabaseConfig()
	
	//Database Migration
	database.ConnectDatabase()
}

func main() {
	// Migrate the schema
	database.DB.AutoMigrate(&models.User{})
}