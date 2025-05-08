package repository

import (
	"context"
	"fmt"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
)

type climbRepository struct {
	queryer Queryer
}

func (s *climbRepository) GetAll(ctx context.Context, mountainID int, categoryID int) ([]domain.Climb, error) {
	const op = "climbRepository.GetAll"

	climbs := make([]domain.Climb, 0)

	query := fmt.Sprintf("SELECT * FROM %s WHERE 1=1", mountain_climbsTable)
	args := []interface{}{}
	argIdx := 1

	if mountainID != 0 {
		query += fmt.Sprintf(" AND id_mountain = $%d", argIdx)
		args = append(args, mountainID)
		argIdx++
	}

	if categoryID != 0 {
		query += fmt.Sprintf(" AND id_category = $%d", argIdx)
		args = append(args, categoryID)
		argIdx++
	}

	if err := s.queryer.SelectContext(ctx, &climbs, query, args...); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return climbs, nil
}

func (s *climbRepository) GetById(ctx context.Context, climbID int64) (domain.Climb, error) {
	climb := domain.Climb{}

	return climb, nil
}
