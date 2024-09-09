package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/verlinof/softlancer-go/config/app_config"
	"github.com/verlinof/softlancer-go/config/db_config"
	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/seeders"
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
	// Seed Tables
	err := seeders.SeedUsers()
	if err != nil {
		log.Fatal("Failed to seed users: ", err)
	}

	err = seeders.SeedRoles()
	if err != nil {
		log.Fatal("Failed to seed roles: ", err)
	}

	err = seeders.SeedReferences()
	if err != nil {
		log.Fatal("Failed to seed references: ", err)
	}

	log.Println("All seeding completed successfully")
}
