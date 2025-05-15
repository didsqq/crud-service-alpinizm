package repository

import (
	"context"
	"fmt"

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

	query := fmt.Sprintf("INSERT INTO %s (alpinist_id, equipment_id, status) VALUES ($1, $2, $3)", alpinistEquipmentTable)

	if _, err := s.queryer.ExecContext(ctx, query, alpinistID, equipmentID, "забронировано"); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	query = fmt.Sprintf("UPDATE %s SET quantity_available = quantity_available - 1 WHERE id = $1", equipmentTable)

	if _, err := s.queryer.ExecContext(ctx, query, equipmentID); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *equipmentRepository) GetAlpinistEquipment(ctx context.Context, alpinistID int64) ([]domain.AlpinistEquipment, error) {
	const op = "equipmentRepository.GetAlpinistEquipment"

	query := fmt.Sprintf(`
		SELECT e.id, e.title, e.quantity_available, e.image_url, e.description, ae.date_of_issue, ae.date_of_return, ae.status
		FROM %s ae
		JOIN %s e ON ae.equipment_id = e.id
		WHERE ae.alpinist_id = $1`,
		alpinistEquipmentTable, equipmentTable)

	var equipments []domain.AlpinistEquipment
	if err := s.queryer.SelectContext(ctx, &equipments, query, alpinistID); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return equipments, nil
}

func (s *equipmentRepository) DeleteAlpinistEquipment(ctx context.Context, alpinistID int64, equipmentID int64) error {
	const op = "equipmentRepository.DeleteAlpinistEquipment"

	query := fmt.Sprintf("UPDATE %s SET status = 'пользователь отменил бронь' WHERE alpinist_id = $1 AND equipment_id = $2", alpinistEquipmentTable)

	if _, err := s.queryer.ExecContext(ctx, query, alpinistID, equipmentID); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *equipmentRepository) UpdateAlpinistEquipment(ctx context.Context, alpinistID int64, equipmentID int64, equipment domain.AlpinistEquipment) error {
	const op = "equipmentRepository.UpdateAlpinistEquipment"

	query := fmt.Sprintf("UPDATE %s SET status = $1, date_of_issue = $2, date_of_return = $3 WHERE alpinist_id = $4 AND equipment_id = $5", alpinistEquipmentTable)

	if _, err := s.queryer.ExecContext(ctx, query, equipment.Status, equipment.DateOfIssue, equipment.DateOfReturn, alpinistID, equipmentID); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *equipmentRepository) GetAllEquipmentAdmin(ctx context.Context) ([]domain.AlpinistsEquipments, error) {
	const op = "equipmentRepository.GetAllEquipmentAdmin"

	query := fmt.Sprintf(`
		SELECT e.id, ae.alpinist_id, ae.equipment_id, e.title, e.quantity_available, e.image_url, e.description, ae.date_of_issue, ae.date_of_return, ae.status
		FROM %s ae
		JOIN %s e ON ae.equipment_id = e.id`, alpinistEquipmentTable, equipmentTable)

	var equipments []domain.AlpinistsEquipments
	if err := s.queryer.SelectContext(ctx, &equipments, query); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return equipments, nil
}
