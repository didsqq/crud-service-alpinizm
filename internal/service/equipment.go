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

func (s *EquipmentService) RecordAlpinistEquipment(ctx context.Context, alpinistID int64, equipmentID int64) error {
	const op = "EquipmentService.RecordAlpinistEquipment"

	err := s.uow.EquipmentsDb().RecordAlpinistEquipment(ctx, alpinistID, equipmentID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *EquipmentService) GetAlpinistEquipment(ctx context.Context, alpinistID int64) ([]domain.Equipment, error) {
	const op = "EquipmentService.GetAlpinistEquipment"

	equipments, err := s.uow.EquipmentsDb().GetAlpinistEquipment(ctx, alpinistID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return equipments, nil
}

func (s *EquipmentService) GetAll(ctx context.Context) ([]domain.Equipment, error) {
	const op = "EquipmentService.GetAll"

	equipments, err := s.uow.EquipmentsDb().GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return equipments, nil
}
