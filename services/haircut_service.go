package services

import (
	"errors"
	"haircut-api-go-gin-gorm-postgres/models"
	"haircut-api-go-gin-gorm-postgres/repositories"
)

type HaircutService interface {
	GetAllHaircuts() ([]models.Haircut, error)
	GetHaircutByID(id uint) (models.Haircut, error)
	CreateHaircut(haircut models.Haircut) (models.Haircut, error)
	UpdateHaircut(haircut models.Haircut) (models.Haircut, error)
	DeleteHaircut(id uint) error
}

type haircutService struct {
	repository repositories.HaircutRepository
}

// CreateHaircut implements HaircutService.
func (h *haircutService) CreateHaircut(haircut models.Haircut) (models.Haircut, error) {
	if haircut.Name == "" {
		return models.Haircut{}, errors.New("haircut name is required")
	}
	return h.repository.Create(haircut)
}

// DeleteHaircut implements HaircutService.
func (h *haircutService) DeleteHaircut(id uint) error {
	panic("unimplemented")
}

// GetAllHaircuts implements HaircutService.
func (h *haircutService) GetAllHaircuts() ([]models.Haircut, error) {
	return h.repository.FindAll()
}

// GetHaircutByID implements HaircutService.
func (h *haircutService) GetHaircutByID(id uint) (models.Haircut, error) {
	return h.repository.FindByID(id)
}

// UpdateHaircut implements HaircutService.
func (h *haircutService) UpdateHaircut(haircut models.Haircut) (models.Haircut, error) {
	panic("unimplemented")
}

func NewHaircutService(repository repositories.HaircutRepository) HaircutService {
	return &haircutService{repository}
}
