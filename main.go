package main

import (
	"boilerplate-webserver-code/src/config"
	"boilerplate-webserver-code/src/logger"
	"boilerplate-webserver-code/src/router"
)

func main() {
	config.Init()
	config.AppConfig = config.GetConfig()

	logger.Init()
	logger.InfoLn("Logger initialized successfully")
	router.Init()
	logger.InfoLn("Router initialized successfully")
}
