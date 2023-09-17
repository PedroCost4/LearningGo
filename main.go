package main

import (
	"github.com/PedroCost4/LearningGo/config"
	"github.com/PedroCost4/LearningGo/router"
)

var (
	logger *config.Logger
)

func main() {

	logger := *config.GetLogger("main")
	//Intitialize config
	err := config.Init()

	if err != nil {
		logger.Errorf("config Initialization error: %v", err)
		return
	}

	//Initialize router
	router.Init()

}
