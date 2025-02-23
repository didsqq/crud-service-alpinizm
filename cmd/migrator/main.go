package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlserver"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	var migrationsPath, direction string

	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&direction, "direction", "up", "migration direction (up or down)")
	flag.Parse()

	if migrationsPath == "" {
		log.Fatal("migrations-path must be specified")
	}

	if direction != "up" && direction != "down" {
		log.Fatal("Ошибка: допустимые значения для direction - 'up' или 'down'")
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	pass := loadEnv("SA_PASSWORD")
	dbName := loadEnv("MSSQL_DATABASE")
	port := loadEnv("PORT")
	login := loadEnv("DB_USERNAME")
	host := loadEnv("DB_HOST")

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable", login, pass, host, port, dbName),
	)
	if err != nil {
		log.Fatalf("migrations create error: %s", err)
	}

	if direction == "up" {
		if err := m.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				fmt.Println("Миграции уже применены. Изменений нет.")
				return
			}
			log.Fatal("Ошибка применения миграций:", err)
		}
		fmt.Println("Миграции успешно выполнены")
	} else if direction == "down" {
		if err := m.Down(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				fmt.Println("Миграции уже откатились. Изменений нет.")
				return
			}
			log.Fatal("Ошибка отката миграций:", err)
		}
		fmt.Println("Миграции успешно откатились")
	}
}

func loadEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("Ошибка: переменная окружения %s не найдена", name)
	}
	return value
}
