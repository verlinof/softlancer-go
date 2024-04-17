package app_config

import "os"

var PORT = ":8000"
var STATIC_PATH = "/public"
var STATIC_DIR = "./public"

func InitAppConfig() {
	portEnv := os.Getenv("APP_PORT")
	if(portEnv != "") {
		PORT = portEnv
	}

	staticPathEnv := os.Getenv("APP_STATIC_PATH")
	if(staticPathEnv != "") {
		STATIC_PATH = staticPathEnv
	}

	staticDirEnv := os.Getenv("APP_STATIC_DIR")
	if(staticDirEnv != "") {
		STATIC_DIR = staticDirEnv
	}
}