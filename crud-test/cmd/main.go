package main

import (
	"fmt"

	"github.com/macreai/arda-ardiyansyah-backend-test/internal/config"
)

func main() {

	viperConfig := config.NewViper("./")
	fiberApp := config.NewFiber(viperConfig)
	logConfig := config.NewLogrus(viperConfig)
	gormConfig := config.NewDatabase(viperConfig, logConfig)
	validatorConfig := config.NewValidator()

	config.InitApp(&config.AppConfig{
		App:       fiberApp,
		DB:        gormConfig,
		Log:       logConfig,
		Validator: validatorConfig,
		Viper:     viperConfig,
	})

	webPort := viperConfig.GetInt("WEB_PORT")
	err := fiberApp.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		logConfig.Fatalf("Failed to start the server : %v", err)
	}
}
