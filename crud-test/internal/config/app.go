package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/macreai/arda-ardiyansyah-backend-test/internal/controller"
	"github.com/macreai/arda-ardiyansyah-backend-test/internal/http"
	"github.com/macreai/arda-ardiyansyah-backend-test/internal/repository"
	"github.com/macreai/arda-ardiyansyah-backend-test/internal/usecase"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AppConfig struct {
	App       *fiber.App
	DB        *gorm.DB
	Log       *logrus.Logger
	Validator *validator.Validate
	Viper     *viper.Viper
}

func InitApp(app *AppConfig) {
	muridRepository := repository.NewMuridRepositoryImpl(app.Log)
	muridUsecase := usecase.NewMuridUseCase(app.DB, app.Log, app.Validator, muridRepository)
	muridController := controller.NewMuridController(app.Log, muridUsecase)

	routeConfig := &http.RouteConfig{
		App:             app.App,
		MuridController: muridController,
	}
	routeConfig.Setup()
}
