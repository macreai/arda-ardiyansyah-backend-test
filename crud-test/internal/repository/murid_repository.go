package repository

import (
	"github.com/macreai/arda-ardiyansyah-backend-test/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MuridRepository interface {
	RegisterMurid(db *gorm.DB, murid *entity.Murid) error
	GetAllMurid(db *gorm.DB) ([]*entity.Murid, error)
}

type MuridRepositoryImpl struct {
	Repository[entity.Murid]
	log *logrus.Logger
}

func NewMuridRepositoryImpl(log *logrus.Logger) *MuridRepositoryImpl {
	return &MuridRepositoryImpl{
		log: log,
	}
}

func (m *MuridRepositoryImpl) RegisterMurid(db *gorm.DB, murid *entity.Murid) error {
	return m.Repository.Create(db, murid)
}

func (m *MuridRepositoryImpl) GetAllMurid(db *gorm.DB) ([]*entity.Murid, error) {
	return m.Repository.GetAll(db)
}
