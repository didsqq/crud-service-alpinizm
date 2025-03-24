package storage

import (
	"context"
	"fmt"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
	"github.com/jmoiron/sqlx"
)

type ClimbStorage struct {
	db *sqlx.DB
}

func (s *ClimbStorage) GetAll(ctx context.Context) ([]domain.Climb, error) {
	const op = "ClimbStorage.GetAll"

	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer conn.Close()

	climbs := make([]domain.Climb, 0)
	query := fmt.Sprintf("SELECT * FROM %s", mountain_climbs)

	err = conn.SelectContext(ctx, &climbs, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return climbs, nil
}

func (s *ClimbStorage) GetById(ctx context.Context, climbID int64) (domain.Climb, error) {
	climb := domain.Climb{}

	return climb, nil
}

func NewClimb(db *sqlx.DB) *ClimbStorage {
	return &ClimbStorage{
		db: db,
	}
}
