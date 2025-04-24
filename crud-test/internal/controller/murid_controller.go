package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/macreai/arda-ardiyansyah-backend-test/internal/model"
	"github.com/macreai/arda-ardiyansyah-backend-test/internal/usecase"
	"github.com/sirupsen/logrus"
)

type MuridController struct {
	Log     *logrus.Logger
	UseCase *usecase.MuridUseCase
}

func NewMuridController(log *logrus.Logger, usecase *usecase.MuridUseCase) *MuridController {
	return &MuridController{
		Log:     log,
		UseCase: usecase,
	}
}

func (m *MuridController) Register(ctx *fiber.Ctx) error {
	request := new(model.RegisterMuridRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		m.Log.Warnf("Failed to parse request body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[*model.RegisterMuridResponse]{
			Errors: fiber.NewError(fiber.ErrBadRequest.Code, fmt.Sprintf("Failed to parse request body: %+v", err)),
			Data:   nil,
		})
	}

	response := m.UseCase.Register(request)

	return ctx.Status(response.Status).JSON(response)
}

func (m *MuridController) GetAllMurid(ctx *fiber.Ctx) error {
	response := m.UseCase.GetAllMurid()

	return ctx.Status(response.Status).JSON(response)
}
