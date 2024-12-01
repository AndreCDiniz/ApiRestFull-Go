package main

import (
	"github.com/AndreCDiniz/ApiRestFull-Go/config"
	"github.com/AndreCDiniz/ApiRestFull-Go/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize Configs
	error := config.Init()
	if error != nil {
		logger.Errorf("Config initialization error: %v", error)
		return
	}

	// Initialize Router
	router.Initialize()

}
