package service

import (
	"context"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
	"github.com/didsqq/crud-service-alpinizm/internal/repository"
)

type Climbs interface {
	GetAll(ctx context.Context) ([]domain.Climb, error)
	GetById(climbID int64) (domain.Climb, error)
}

type User interface {
	Create(ctx context.Context, user domain.User) (int, error)
	GetByID(ctx context.Context, id int) (*domain.User, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]domain.User, error)
}

type Service struct {
	User
	Climbs
}

func NewService(repo repository.UnitOfWork) *Service {
	return &Service{
		User:   NewUserService(repo),
		Climbs: NewClimbsService(repo),
	}
}
