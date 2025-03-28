package storage

import (
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

const (
	category_of_difficulty = "category_of_difficulty"
	sport_category         = "sport_category"
	position               = "position"
	alpinists              = "alpinists"
	equipment              = "equipment"
	mountain               = "mountain"
	groups                 = "groups"
	mountain_climbs        = "mountain_climbs"
	equipment_inventory    = "equipment_inventory"
	team                   = "team"
	team_leaders           = "team_leaders"
	climbers_in_groups     = "climbers_in_groups"
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
