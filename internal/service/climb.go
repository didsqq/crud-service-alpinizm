package service

import (
	"context"
	"fmt"
	"log"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
	"github.com/didsqq/crud-service-alpinizm/internal/repository"
)

type ClimbService struct {
	uow repository.UnitOfWork
}

func NewClimbsService(uow repository.UnitOfWork) *ClimbService {
	return &ClimbService{
		uow: uow,
	}
}

func (s *ClimbService) RecordAlpinistClimb(ctx context.Context, alpinistID int64, climbID int64) error {
	const op = "ClimbService.RecordAlpinistClimb"

	if err := s.uow.Begin(); err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	var transactionErr error
	defer func() {
		if transactionErr != nil {
			if rbErr := s.uow.Rollback(); rbErr != nil {
				log.Printf("failed to rollback transaction: %v", rbErr)
			}
		}
	}()

	err := s.uow.ClimbsTx().CheckRecordAlpinistClimb(ctx, alpinistID, climbID)
	if err != nil {
		transactionErr = fmt.Errorf("%s: %w", op, err)
		return transactionErr
	}

	err = s.uow.ClimbsTx().RecordAlpinistClimb(ctx, alpinistID, climbID)
	if err != nil {
		transactionErr = fmt.Errorf("%s: %w", op, err)
		return transactionErr
	}

	if err := s.uow.Commit(); err != nil {
		transactionErr = fmt.Errorf("failed to commit transaction: %w", err)
		return transactionErr
	}

	return nil
}

func (s *ClimbService) GetAll(ctx context.Context, mountainID int, categoryID int) ([]domain.Climb, error) {
	const op = "ClimbService.GetAll"

	climbs, err := s.uow.ClimbsDb().GetAll(ctx, mountainID, categoryID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return climbs, nil
}

func (s *ClimbService) GetById(ctx context.Context, climbID int64) (domain.Climb, error) {
	const op = "ClimbService.GetById"

	climb, err := s.uow.ClimbsDb().GetById(ctx, climbID)
	if err != nil {
		return domain.Climb{}, fmt.Errorf("%s: %w", op, err)
	}

	return climb, nil
}
