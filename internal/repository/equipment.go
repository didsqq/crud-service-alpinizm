package repository

import (
	"context"
	"fmt"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
)

type equipmentRepository struct {
	queryer Queryer
}

func (s *equipmentRepository) GetAll(ctx context.Context) ([]domain.Equipment, error) {
	const op = "equipmentRepository.GetAll"

	equipments := make([]domain.Equipment, 0)
	query := fmt.Sprintf("SELECT * FROM %s", equipmentTable)

	if err := s.queryer.SelectContext(ctx, &equipments, query); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return equipments, nil
}
