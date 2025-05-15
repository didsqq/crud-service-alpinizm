package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
	"github.com/jmoiron/sqlx"
)

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrUserNameExist         = errors.New("username exist")
	ErrAlpinistHasRegistered = errors.New("альпинист уже зарегистрирован на это восхождение")
)

type ClimbRepository interface {
	GetAll(ctx context.Context, mountainID int, categoryID int) ([]domain.Climb, error)
	GetById(ctx context.Context, climbID int64) (domain.Climb, error)
	RecordAlpinistClimb(ctx context.Context, alpinistID int64, climbID int64) error
	CheckRecordAlpinistClimb(ctx context.Context, alpinistID int64, climbID int64) error
	GetAlpinistClimb(ctx context.Context, alpinistID int64) ([]domain.Climb, error)
	GetAllCategoryOfDifficulty(ctx context.Context) ([]domain.CategoryOfDifficulty, error)
}

type UserRepository interface {
	Create(ctx context.Context, user domain.User) (int, error)
	GetByID(ctx context.Context, id int) (*domain.User, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]domain.User, error)
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	GetAllSportCategory(ctx context.Context) ([]domain.SportCategory, error)
	GetAllAlpinists(ctx context.Context) ([]domain.User, error)
	CancelAlpinistClimb(ctx context.Context, alpinistID int64, climbID int64) error
}

type EquipmentRepository interface {
	GetAll(ctx context.Context) ([]domain.Equipment, error)
	RecordAlpinistEquipment(ctx context.Context, alpinistID int64, equipmentID int64) error
	GetAlpinistEquipment(ctx context.Context, alpinistID int64) ([]domain.AlpinistEquipment, error)
	DeleteAlpinistEquipment(ctx context.Context, alpinistID int64, equipmentID int64) error
}

type MountainRepository interface {
	GetAll(ctx context.Context) ([]domain.Mountain, error)
}

type UnitOfWork interface {
	Begin() error
	Commit() error
	Rollback() error
	UsersDb() UserRepository
	ClimbsDb() ClimbRepository
	EquipmentsDb() EquipmentRepository
	MountainsDb() MountainRepository
	ClimbsTx() ClimbRepository
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

func (u *unitOfWork) ClimbsTx() ClimbRepository {
	return &climbRepository{queryer: u.tx}
}

func (u *unitOfWork) EquipmentsDb() EquipmentRepository {
	return &equipmentRepository{queryer: u.db}
}

func (u *unitOfWork) MountainsDb() MountainRepository {
	return &mountainRepository{queryer: u.db}
}
