package service

import (
	"context"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
	"github.com/didsqq/crud-service-alpinizm/internal/repository"
	"github.com/go-chi/jwtauth"
)

type Climbs interface {
	GetAll(ctx context.Context, mountainID int, categoryID int) ([]domain.Climb, error)
	GetById(ctx context.Context, climbID int64) (domain.Climb, error)
}

type User interface {
	Create(ctx context.Context, user domain.User) (int, error)
	GetByID(ctx context.Context, id int) (*domain.User, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]domain.User, error)
	Login(ctx context.Context, username string, password string) (string, error)
	GetAllSportCategory(ctx context.Context) ([]domain.SportCategory, error)
	CheckToken(ctx context.Context, token string) (bool, error)
}

type Equipments interface {
	GetAll(ctx context.Context) ([]domain.Equipment, error)
}

type Mountains interface {
	GetAll(ctx context.Context) ([]domain.Mountain, error)
}

type Service struct {
	User
	Climbs
	Equipments
	Mountains
}

func NewService(repo repository.UnitOfWork, tokenAuth *jwtauth.JWTAuth) *Service {
	return &Service{
		User:       NewUserService(repo, tokenAuth),
		Climbs:     NewClimbsService(repo),
		Equipments: NewEquipmentService(repo),
		Mountains:  NewMountainService(repo),
	}
}
