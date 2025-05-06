package service

import (
	"context"
	"fmt"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
	"github.com/didsqq/crud-service-alpinizm/internal/repository"
)

type EquipmentService struct {
	uow repository.UnitOfWork
}

func NewEquipmentService(uow repository.UnitOfWork) *EquipmentService {
	return &EquipmentService{
		uow: uow,
	}
}

func (s *EquipmentService) GetAll(ctx context.Context) ([]domain.Equipment, error) {
	const op = "EquipmentService.GetAll"

	equipments, err := s.uow.EquipmentsDb().GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return equipments, nil
}
