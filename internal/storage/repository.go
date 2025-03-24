package storage

import (
	"context"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Climbs interface {
	GetAll(ctx context.Context) ([]domain.Climb, error)
	GetById(ctx context.Context, climbID int64) (domain.Climb, error)
}

type Repository struct {
	Climbs
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Climbs: NewClimb(db),
	}
}