package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	category_of_difficultyTable = "category_of_difficulty"
	sport_categoryTable         = "sport_category"
	positionTable               = "position"
	alpinistsTable              = "alpinists"
	equipmentTable              = "equipment"
	mountainTable               = "mountain"
	groupsTable                 = "groups"
	mountain_climbsTable        = "mountain_climbs"
	equipment_inventoryTable    = "equipment_inventory"
	teamTable                   = "team"
	team_leadersTable           = "team_leaders"
	climbers_in_groupsTable     = "climbers_in_groups"
	alpinistEquipmentTable      = "alpinist_equipment"
)

func NewPostgresDB(dsn string) (*sqlx.DB, error) {
	const op = "repository.NewPostgresDB"

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return db, nil
}
