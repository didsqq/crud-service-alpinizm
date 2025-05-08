package repository

import (
	"context"
	"fmt"
	"github.com/didsqq/crud-service-alpinizm/internal/domain"
)

type mountainRepository struct {
	queryer Queryer
}

func (s *mountainRepository) GetAll(ctx context.Context) ([]domain.Mountain, error) {
	const op = "mountainRepository.GetAll"

	mountains := make([]domain.Mountain, 0)

	query := fmt.Sprintf("SELECT * FROM %s", mountainTable)

	if err := s.queryer.SelectContext(ctx, &mountains, query); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return mountains, nil
}
