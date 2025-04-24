package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/macreai/arda-ardiyansyah-backend-test/internal/controller"
)

type RouteConfig struct {
	App             *fiber.App
	MuridController *controller.MuridController
}

func (r *RouteConfig) Setup() {

	r.App.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,PATCH,DELETE",
	}))

	r.setupGuestRoute()
}

func (r *RouteConfig) setupGuestRoute() {
	r.App.Post("/api/v1/users/register", r.MuridController.Register)
	r.App.Get("/api/v1/murids", r.MuridController.GetAllMurid)
}
