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

func (r *userRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	const op = "userRepository.GetAll"

	var users []domain.User

	return users, nil
}

func (r *userRepository) Create(ctx context.Context, user domain.User) (int, error) {
	const op = "UserRepository.Create"

	query := fmt.Sprintf(`
		INSERT INTO %s 
		(Surname, Name_, Address_, Phone, Sex, ID_sport_category, Username, Password_)
		OUTPUT INSERTED.ID_alpinist
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`, alpinistsTable)

	row := r.queryer.QueryRowxContext(ctx, query,
		user.Surname,
		user.Name,
		user.Address,
		user.Phone,
		user.Sex,
		user.ID_sport_category,
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
