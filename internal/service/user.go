package service

import (
	"context"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
	"github.com/didsqq/crud-service-alpinizm/internal/repository"
)

type UserService struct {
	uow repository.UnitOfWork
}

func NewUserService(uow repository.UnitOfWork) *UserService {
	return &UserService{
		uow: uow,
	}
}

func (s *UserService) GetAll(ctx context.Context) ([]domain.User, error) {
	users, err := s.uow.UsersDb().GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) Create(ctx context.Context, user domain.User) (int, error) {
	id, err := s.uow.UsersDb().Create(ctx, user)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *UserService) GetByID(ctx context.Context, id int) (*domain.User, error) {
	user, err := s.uow.UsersDb().GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// books, err := s.uow.BooksDb().GetByUserID(ctx, id)
	// if err != nil {
	// 	return nil, err
	// }

	// user.RentedBooks = *books

	return user, nil
}

func (s *UserService) Delete(ctx context.Context, id int) error {
	if err := s.uow.UsersDb().Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
