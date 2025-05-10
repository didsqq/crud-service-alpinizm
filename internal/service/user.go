package service

import (
	"context"
	"time"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
	"github.com/didsqq/crud-service-alpinizm/internal/repository"
	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	uow       repository.UnitOfWork
	tokenAuth *jwtauth.JWTAuth
}

func NewUserService(uow repository.UnitOfWork, tokenAuth *jwtauth.JWTAuth) *UserService {
	return &UserService{
		uow:       uow,
		tokenAuth: tokenAuth,
	}
}

func (s *UserService) CheckToken(ctx context.Context, token string) (bool, error) {
	_, err := s.tokenAuth.Decode(token)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *UserService) GetAllSportCategory(ctx context.Context) ([]domain.SportCategory, error) {
	users, err := s.uow.UsersDb().GetAllSportCategory(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) GetAll(ctx context.Context) ([]domain.User, error) {
	users, err := s.uow.UsersDb().GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) Login(ctx context.Context, username string, password string) (string, error) {
	user, err := s.uow.UsersDb().GetByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	_, token, err := s.tokenAuth.Encode(map[string]interface{}{
		"id":    user.ID,
		"login": user.Username,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})
	if err != nil {
		return "", err
	}

	return token, nil

}

func (s *UserService) Create(ctx context.Context, user domain.User) (int, error) {
	passHash, err := generateHash(user.Password)
	if err != nil {
		return 0, err
	}

	user.Password = passHash

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

	return user, nil
}

func (s *UserService) Delete(ctx context.Context, id int) error {
	if err := s.uow.UsersDb().Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func generateHash(pass string) (string, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(passHash), nil
}
