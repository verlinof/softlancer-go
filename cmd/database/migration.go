package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/verlinof/softlancer-go/config/app_config"
	"github.com/verlinof/softlancer-go/config/db_config"
	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/models"
)

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println(err)
	}

	// Init config
	app_config.InitAppConfig()
	db_config.InitDatabaseConfig()

	// Database connection
	database.ConnectDatabase()
}

func main() {
	// Drop the tables if they exist
	err := database.DB.Migrator().DropTable(&models.User{}, &models.Company{}, &models.Project{}, &models.Role{}, &models.Reference{}, &models.Application{})
	if err != nil {
		log.Println("Error dropping tables: ", err.Error())
	} else {
		log.Println("Tables dropped successfully.")
	}

	// Migrate the schema
	err = database.DB.AutoMigrate(&models.User{}, &models.Company{}, &models.Project{}, &models.Role{}, &models.Reference{}, &models.Application{})
	if err != nil {
		log.Println("Error migrating schema: ", err.Error())
	} else {
		log.Println("Migration complete!")
	}
}
