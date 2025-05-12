package repository

import (
	"context"
	"fmt"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
)

type climbRepository struct {
	queryer Queryer
}

func (s *climbRepository) CheckRecordAlpinistClimb(ctx context.Context, alpinistID int64, climbID int64) error {
	const op = "climbRepository.CheckRecordAlpinistClimb"

	query := `
	SELECT COUNT(*) FROM alpinist_climb WHERE alpinist_id = $1 AND climb_id = $2
	`
	var count int
	err := s.queryer.GetContext(ctx, &count, query, alpinistID, climbID)
	if err != nil {
		return fmt.Errorf("%s: failed to check record alpinist climb: %w", op, err)
	}

	if count > 0 {
		return ErrAlpinistHasRegistered
	}

	return nil
}

func (s *climbRepository) RecordAlpinistClimb(ctx context.Context, alpinistID int64, climbID int64) error {
	const op = "climbRepository.RecordAlpinistClimb"

	query := `
	INSERT INTO alpinist_climb (alpinist_id, climb_id)
	VALUES ($1, $2)
	`
	_, err := s.queryer.ExecContext(ctx, query, alpinistID, climbID)
	if err != nil {
		return fmt.Errorf("%s: failed to record alpinist climb: %w", op, err)
	}

	query = `
	UPDATE mountain_climbs
	SET places_left = places_left - 1
	WHERE id = $1
	`

	_, err = s.queryer.ExecContext(ctx, query, climbID)
	if err != nil {
		return fmt.Errorf("%s: failed to update climb places left: %w", op, err)
	}

	return nil
}

func (s *climbRepository) GetAll(ctx context.Context, mountainID int, categoryID int) ([]domain.Climb, error) {
	const op = "climbRepository.GetAll"

	climbs := make([]domain.Climb, 0)

	query := fmt.Sprintf("SELECT * FROM %s WHERE 1=1", mountain_climbsTable)
	args := []interface{}{}
	argIdx := 1

	if mountainID != 0 {
		query += fmt.Sprintf(" AND id_mountain = $%d", argIdx)
		args = append(args, mountainID)
		argIdx++
	}

	if categoryID != 0 {
		query += fmt.Sprintf(" AND id_category = $%d", argIdx)
		args = append(args, categoryID)
		argIdx++
	}

	if err := s.queryer.SelectContext(ctx, &climbs, query, args...); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return climbs, nil
}

func (s *climbRepository) GetById(ctx context.Context, climbID int64) (domain.Climb, error) {
	const op = "climbRepository.GetById"

	var climb domain.Climb

	// 1. Основные данные
	query := `
	SELECT 
		id, id_mountain, id_category, title, season, duration,
		distance, elevation, map_url, rating, description, 
		start_date, end_date, total, places_left, photo_url
	FROM mountain_climbs
	WHERE id = $1
	`
	err := s.queryer.GetContext(ctx, &climb, query, climbID)
	if err != nil {
		return domain.Climb{}, fmt.Errorf("%s: failed to get climb: %w", op, err)
	}

	// 2. Руководители
	leadersQuery := `
	SELECT t.id, t.surname_name, t.experience, t.date_of_birth, t.address_, t.id_position, t.phone, t.password_, t.login_
	FROM team_leaders tl
	JOIN team t ON t.id = tl.id_team_member
	WHERE tl.id_mountain_climb = $1
	`
	err = s.queryer.SelectContext(ctx, &climb.TeamLeaders, leadersQuery, climbID)
	if err != nil {
		return domain.Climb{}, fmt.Errorf("%s: failed to get team leaders: %w", op, err)
	}

	// 3. Снаряжение
	equipmentQuery := `
	SELECT e.id, e.title, e.quantity_available, e.image_url, e.description
	FROM climb_equipment ce
	JOIN equipment e ON e.id = ce.equipment_id
	WHERE ce.climb_id = $1
	`
	err = s.queryer.SelectContext(ctx, &climb.Equipments, equipmentQuery, climbID)
	if err != nil {
		return domain.Climb{}, fmt.Errorf("%s: failed to get equipment: %w", op, err)
	}

	// 4. Изображения
	imagesQuery := `
	SELECT id, climb_id, url
	FROM climb_images
	WHERE climb_id = $1
	`
	err = s.queryer.SelectContext(ctx, &climb.Images, imagesQuery, climbID)
	if err != nil {
		return domain.Climb{}, fmt.Errorf("%s: failed to get images: %w", op, err)
	}

	// 5. Гора
	mountainQuery := `
	SELECT id, title, height, mountain_range
	FROM mountain
	WHERE id = $1
	`
	err = s.queryer.GetContext(ctx, &climb.Mountain, mountainQuery, climb.IdMountain)
	if err != nil {
		return domain.Climb{}, fmt.Errorf("%s: failed to get mountain: %w", op, err)
	}

	// 6. Категория
	categoryQuery := `
	SELECT title
	FROM sport_category
	WHERE id = $1
	`
	err = s.queryer.GetContext(ctx, &climb.Category, categoryQuery, climb.IdMountain)
	if err != nil {
		return domain.Climb{}, fmt.Errorf("%s: failed to get category: %w", op, err)
	}

	return climb, nil
}
