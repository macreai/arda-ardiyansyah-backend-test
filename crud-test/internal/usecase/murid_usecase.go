package usecase

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/macreai/arda-ardiyansyah-backend-test/internal/entity"
	"github.com/macreai/arda-ardiyansyah-backend-test/internal/model"
	"github.com/macreai/arda-ardiyansyah-backend-test/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MuridUseCase struct {
	DB              *gorm.DB
	Log             *logrus.Logger
	Validate        *validator.Validate
	MuridRepository *repository.MuridRepositoryImpl
}

func NewMuridUseCase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	muridRepository *repository.MuridRepositoryImpl,
) *MuridUseCase {
	return &MuridUseCase{
		DB:              db,
		Log:             logger,
		Validate:        validate,
		MuridRepository: muridRepository,
	}
}

func (m *MuridUseCase) Register(request *model.RegisterMuridRequest) *model.WebResponse[*model.RegisterMuridResponse] {
	err := m.Validate.Struct(request)
	if err != nil {
		m.Log.Warnf("Invalid request body: %+v", err)
		return &model.WebResponse[*model.RegisterMuridResponse]{
			Errors: fiber.NewError(fiber.ErrBadRequest.Code, fmt.Sprintf("Invalid request body: %+v", err)),
			Data:   nil,
			Status: fiber.StatusBadRequest,
		}
	}

	user := &entity.Murid{
		Name:       request.Name,
		TimeCreate: time.Now(),
	}

	if err := m.MuridRepository.Create(m.DB, user); err != nil {
		return &model.WebResponse[*model.RegisterMuridResponse]{
			Errors: fiber.NewError(fiber.ErrInternalServerError.Code, err.Error()),
			Data:   nil,
			Status: fiber.StatusInternalServerError,
		}

	}

	return &model.WebResponse[*model.RegisterMuridResponse]{
		Errors: nil,
		Data: &model.RegisterMuridResponse{
			Message: "Success create murid",
		},
		Status: fiber.StatusOK,
	}

}

func (m *MuridUseCase) GetAllMurid() *model.WebResponse[*model.GetAllMuridResponse] {

	if murids, err := m.MuridRepository.GetAllMurid(m.DB); err != nil {
		return &model.WebResponse[*model.GetAllMuridResponse]{
			Errors: fiber.NewError(fiber.ErrInternalServerError.Code, err.Error()),
			Data:   nil,
			Status: fiber.StatusInternalServerError,
		}
	} else {
		return &model.WebResponse[*model.GetAllMuridResponse]{
			Errors: nil,
			Data: &model.GetAllMuridResponse{
				Murids: murids,
			},
			Status: fiber.StatusOK,
		}
	}
}
