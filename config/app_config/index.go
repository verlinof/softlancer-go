package app_config

import "os"

var GIN_MODE = "debug"
var PORT = ":8000"
var BASE_URL = "http://localhost:8000"
var APP_DIR = "./"
var STATIC_PATH = "/storage"
var STATIC_DIR = "./storage"

func Init() {
	ginModeEnv := os.Getenv("GIN_MODE")
	if ginModeEnv != "" {
		GIN_MODE = ginModeEnv
	}

	portEnv := os.Getenv("APP_PORT")
	if portEnv != "" {
		PORT = portEnv
	}

	staticPathEnv := os.Getenv("APP_STATIC_PATH")
	if staticPathEnv != "" {
		STATIC_PATH = staticPathEnv
	}

	appDirEnv := os.Getenv("APP_DIR")
	if appDirEnv != "" {
		APP_DIR = appDirEnv
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
