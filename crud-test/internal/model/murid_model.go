package model

import "github.com/macreai/arda-ardiyansyah-backend-test/internal/entity"

type RegisterMuridRequest struct {
	Name string `json:"name" validate:"required"`
}

type RegisterMuridResponse struct {
	Message string `json:"message"`
}

type GetAllMuridResponse struct {
	Murids []*entity.Murid `json:"murids"`
}
