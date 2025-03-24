package services

import (
	"context"
	"fmt"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
	"github.com/didsqq/crud-service-alpinizm/internal/storage"
)

type ClimbService struct {
	climbStorage storage.Climbs
}

func (s *ClimbService) GetAll() ([]domain.Climb, error) {
	const op = "ClimbService.GetAll"

	ctx := context.Background()
	climbs, err := s.climbStorage.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return climbs, nil
}

func (s *ClimbService) GetById(climbID int64) (domain.Climb, error) {
	climb := domain.Climb{}

	return climb, nil
}

func NewClimbsService(repo storage.Climbs) *ClimbService {
	return &ClimbService{
		climbStorage: repo,
	}
}
