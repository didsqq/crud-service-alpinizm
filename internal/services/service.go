package services

import (
	"github.com/didsqq/crud-service-alpinizm/internal/domain"
	"github.com/didsqq/crud-service-alpinizm/internal/storage"
)

type Climbs interface {
	GetAll() ([]domain.Climb, error)
	GetById(climbID int64) (domain.Climb, error)
}

type Service struct {
	Climbs
}

func NewService(repo *storage.Repository) *Service {
	return &Service{
		Climbs: NewClimbsService(repo.Climbs),
	}
}
