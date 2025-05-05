package repository

import (
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
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
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	Encrypt  string
}

func NewMsSqlDB(cfg Config) (*sqlx.DB, error) {
	const op = "storage.mssql.New"
	log.Println(op, cfg)
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.Encrypt)

	db, err := sqlx.Open("sqlserver", connString)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return db, nil
}
