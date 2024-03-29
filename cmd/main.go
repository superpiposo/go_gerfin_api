package main

import (
	"github.com/superpiposo/go_gerfin_api/configs"
	"github.com/superpiposo/go_gerfin_api/pkg/router"
)

func main() {
	err := configs.Init()
	logger := configs.GetLogger("main")

	if err != nil {
		logger.Errorf("Error initializing configs: %s", err)
	}
	router.InitRouter()
}
