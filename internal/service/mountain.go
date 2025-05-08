package service

import (
	"context"
	"fmt"
	"github.com/didsqq/crud-service-alpinizm/internal/domain"
	"github.com/didsqq/crud-service-alpinizm/internal/repository"
)

type MountainService struct {
	uow repository.UnitOfWork
}

func NewMountainService(uow repository.UnitOfWork) *MountainService {
	return &MountainService{
		uow: uow,
	}
}

func (s *MountainService) GetAll(ctx context.Context) ([]domain.Mountain, error) {
	const op = "MountainService.GetAll"

	mountains, err := s.uow.MountainsDb().GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return mountains, nil
}
