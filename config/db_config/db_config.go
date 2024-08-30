package db_config

import "os"

var DB_DRIVER = "mysql"
var DB_HOST = "localhost"
var DB_PORT = "3306"
var DB_USER = "root"
var DB_PASSWORD = ""
var DB_NAME = "softlancer"

func InitDatabaseConfig() {
	driverEnv := os.Getenv("DB_DRIVER")
	if driverEnv != "" {
		DB_DRIVER = driverEnv
	}

	hostEnv := os.Getenv("DB_HOST")
	if hostEnv != "" {
		DB_HOST = hostEnv
	}

	portEnv := os.Getenv("DB_PORT")
	if portEnv != "" {
		DB_PORT = portEnv
	}

	userEnv := os.Getenv("DB_USER")
	if userEnv != "" {
		DB_USER = userEnv
	}

	passwordEnv := os.Getenv("DB_PASSWORD")
	if passwordEnv != "" {
		DB_PASSWORD = passwordEnv
	}

	nameEnv := os.Getenv("DB_NAME")
	if nameEnv != "" {
		DB_NAME = nameEnv
	}
}
