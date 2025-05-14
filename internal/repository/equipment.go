package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
)

type equipmentRepository struct {
	queryer Queryer
}

func (s *equipmentRepository) GetAll(ctx context.Context) ([]domain.Equipment, error) {
	const op = "equipmentRepository.GetAll"

	equipments := make([]domain.Equipment, 0)
	query := fmt.Sprintf("SELECT * FROM %s", equipmentTable)

	if err := s.queryer.SelectContext(ctx, &equipments, query); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return equipments, nil
}

func (s *equipmentRepository) RecordAlpinistEquipment(ctx context.Context, alpinistID int64, equipmentID int64) error {
	const op = "equipmentRepository.RecordAlpinistEquipment"

	query := fmt.Sprintf("INSERT INTO %s (alpinist_id, equipment_id, date_of_issue, date_of_return) VALUES ($1, $2, $3, $4)", alpinistEquipmentTable)

	if _, err := s.queryer.ExecContext(ctx, query, alpinistID, equipmentID, time.Now(), time.Now().Add(time.Hour*24*30)); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	query = fmt.Sprintf("UPDATE %s SET quantity_available = quantity_available - 1 WHERE id = $1", equipmentTable)

	if _, err := s.queryer.ExecContext(ctx, query, equipmentID); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *equipmentRepository) GetAlpinistEquipment(ctx context.Context, alpinistID int64) ([]domain.Equipment, error) {
	const op = "equipmentRepository.GetAlpinistEquipment"

	query := fmt.Sprintf(`
		SELECT e.id, e.title, e.quantity_available, e.image_url, e.description
		FROM %s ae
		JOIN %s e ON ae.equipment_id = e.id
		WHERE ae.alpinist_id = $1`,
		alpinistEquipmentTable, equipmentTable)

	var equipments []domain.Equipment
	if err := s.queryer.SelectContext(ctx, &equipments, query, alpinistID); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return equipments, nil
}

func (s *equipmentRepository) DeleteAlpinistEquipment(ctx context.Context, alpinistID int64, equipmentID int64) error {
	const op = "equipmentRepository.DeleteAlpinistEquipment"

	query := fmt.Sprintf("DELETE FROM %s WHERE alpinist_id = $1 AND equipment_id = $2", alpinistEquipmentTable)

	if _, err := s.queryer.ExecContext(ctx, query, alpinistID, equipmentID); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
