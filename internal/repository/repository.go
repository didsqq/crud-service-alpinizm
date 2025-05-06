package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
	"github.com/jmoiron/sqlx"
)

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrUserNameExist = errors.New("username exist")
)

type ClimbRepository interface {
	GetAll(ctx context.Context) ([]domain.Climb, error)
	GetById(ctx context.Context, climbID int64) (domain.Climb, error)
}

type UserRepository interface {
	Create(ctx context.Context, user domain.User) (int, error)
	GetByID(ctx context.Context, id int) (*domain.User, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]domain.User, error)
}

type EquipmentRepository interface {
	GetAll(ctx context.Context) ([]domain.Equipment, error)
}

type UnitOfWork interface {
	Begin() error
	Commit() error
	Rollback() error
	UsersDb() UserRepository
	ClimbsDb() ClimbRepository
	EquipmentsDb() EquipmentRepository
}

type Queryer interface {
	QueryRowxContext(ctx context.Context, query string, args ...any) *sqlx.Row
	GetContext(ctx context.Context, dest any, query string, args ...any) error
	ExecContext(ctx context.Context, query string, arg ...any) (sql.Result, error)
	SelectContext(ctx context.Context, dest any, query string, args ...any) error
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

type unitOfWork struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

func NewUnitOfWork(db *sqlx.DB) UnitOfWork {
	return &unitOfWork{db: db}
}

func (u *unitOfWork) Begin() error {
	tx, err := u.db.Beginx()
	if err != nil {
		return err
	}
	u.tx = tx
	return nil
}

func (u *unitOfWork) Commit() error {
	return u.tx.Commit()
}

func (u *unitOfWork) Rollback() error {
	return u.tx.Rollback()
}

func (u *unitOfWork) UsersDb() UserRepository {
	return &userRepository{queryer: u.db}
}

func (u *unitOfWork) ClimbsDb() ClimbRepository {
	return &climbRepository{queryer: u.db}
}

func (u *unitOfWork) EquipmentsDb() EquipmentRepository {
	return &equipmentRepository{queryer: u.db}
}
