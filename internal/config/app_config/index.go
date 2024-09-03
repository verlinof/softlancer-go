package app_config

import "os"

var PORT = ":8000"
var BASE_URL = "http://localhost:8000"
var STATIC_PATH = "/storage"
var STATIC_DIR = "./storage"

func InitAppConfig() {
	portEnv := os.Getenv("APP_PORT")
	if portEnv != "" {
		PORT = portEnv
	}

	staticPathEnv := os.Getenv("APP_STATIC_PATH")
	if staticPathEnv != "" {
		STATIC_PATH = staticPathEnv
	}

	staticDirEnv := os.Getenv("APP_STATIC_DIR")
	if staticDirEnv != "" {
		STATIC_DIR = staticDirEnv
	}

	baseUrlEnv := os.Getenv("APP_URL")
	if baseUrlEnv != "" {
		BASE_URL = baseUrlEnv
	}
}
