package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
)

type userRepository struct {
	queryer Queryer
}

func (r *userRepository) GetAllSportCategory(ctx context.Context) ([]domain.SportCategory, error) {
	const op = "userRepository.GetAllSportCategory"

	c := make([]domain.SportCategory, 0)

	query := fmt.Sprintf("SELECT * FROM %s", sport_categoryTable)

	if err := r.queryer.SelectContext(ctx, &c, query); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return c, nil
}

func (r *userRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	const op = "userRepository.GetAll"

	var users []domain.User

	return users, fmt.Errorf("not implemented")
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	const op = "UserRepository.GetByUsername"

	query := fmt.Sprintf(`
		SELECT * FROM %s 
		WHERE username=$1
	`, alpinistsTable)

	row := r.queryer.QueryRowxContext(ctx, query, username)

	user := new(domain.User)
	err := row.Scan(
		&user.ID,
		&user.Surname,
		&user.Name,
		&user.Address,
		&user.Phone,
		&user.Sex,
		&user.IdSportCategory,
		&user.Username,
		&user.Password,
	)
	log.Println(user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (r *userRepository) Create(ctx context.Context, user domain.User) (int, error) {
	const op = "UserRepository.Create"

	query := fmt.Sprintf(`
		INSERT INTO %s 
		(surname, name_, address_, phone, sex, id_sport_category, username, password_)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`, alpinistsTable)

	row := r.queryer.QueryRowxContext(ctx, query,
		user.Surname,
		user.Name,
		user.Address,
		user.Phone,
		user.Sex,
		user.IdSportCategory,
		user.Username,
		user.Password,
	)

	log.Println(user)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("%s: ошибка при создании пользователя: %w", op, err)
	}

	return id, nil
}

func (r *userRepository) GetByID(ctx context.Context, id int) (*domain.User, error) {
	const op = "UserRepository.GetByID"

	query := fmt.Sprintf(`
		SELECT * FROM %s 
		WHERE id=$1
	`, alpinistsTable)

	var user domain.User

	if err := r.queryer.GetContext(ctx, &user, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &user, nil
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
	const op = "UserRepository.DeleteByID"

	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE id = $1
	`, alpinistsTable)

	result, err := r.queryer.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("%s: ошибка при удалении пользователя: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: не удалось получить число удалённых строк: %w", op, err)
	}

	if rowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}

func (r *userRepository) GetAllAlpinists(ctx context.Context) ([]domain.User, error) {
	const op = "userRepository.GetAllAlpinists"

	var users []domain.User

	query := fmt.Sprintf(`
		SELECT * FROM %s
	`, alpinistsTable)

	if err := r.queryer.SelectContext(ctx, &users, query); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return users, nil
}
