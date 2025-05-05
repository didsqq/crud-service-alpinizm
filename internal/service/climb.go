package service

import (
	"context"
	"fmt"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
	"github.com/didsqq/crud-service-alpinizm/internal/repository"
)

type ClimbService struct {
	uow repository.UnitOfWork
}

func NewClimbsService(uow repository.UnitOfWork) *ClimbService {
	return &ClimbService{
		uow: uow,
	}
}

func (s *ClimbService) GetAll(ctx context.Context) ([]domain.Climb, error) {
	const op = "ClimbService.GetAll"

	climbs, err := s.uow.ClimbsDb().GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return climbs, nil
}

func (s *ClimbService) GetById(climbID int64) (domain.Climb, error) {
	climb := domain.Climb{}

	return climb, nil
}
