package repository

import (
	"context"
	"fmt"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
)

type climbRepository struct {
	queryer Queryer
}

func (s *climbRepository) GetAll(ctx context.Context) ([]domain.Climb, error) {
	const op = "ClimbStorage.GetAll"

	climbs := make([]domain.Climb, 0)
	query := fmt.Sprintf("SELECT * FROM %s", mountain_climbsTable)

	if err := s.queryer.SelectContext(ctx, &climbs, query); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return climbs, nil
}

func (s *climbRepository) GetById(ctx context.Context, climbID int64) (domain.Climb, error) {
	climb := domain.Climb{}

	return climb, nil
}
