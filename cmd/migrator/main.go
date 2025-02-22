package main

import (
	"database/sql"
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
	var migrationsPath string

	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.Parse()

	if migrationsPath == "" {
		panic("migrations-path  must be specified")
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	pass := os.Getenv("SA_PASSWORD")
	if pass == "" {
		log.Fatalf("Ошибка: переменная окружения SA_PASSWORD не найдена")
	}
	dbName := os.Getenv("MSSQL_DATABASE")
	if dbName == "" {
		log.Fatalf("Ошибка: переменная окружения MSSQL_DATABASE не найдена")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("Ошибка: переменная окружения PORT не найдена")
	}
	login := "sa"
	host := "localhost"

	masterDBConnStr := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=master", login, pass, host, port)
	db, err := sql.Open("sqlserver", masterDBConnStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к master: %v", err)
	}
	defer db.Close()

	var exists int
	query := fmt.Sprintf("SELECT COUNT(*) FROM sys.databases WHERE name = N'%s'", dbName)
	err = db.QueryRow(query).Scan(&exists)
	if err != nil {
		log.Fatalf("Ошибка при проверке базы: %v", err)
	}

	if exists == 0 {
		fmt.Println("Создаю базу данных...")
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName))
		if err != nil {
			log.Fatalf("Ошибка при создании базы: %v", err)
		}
		fmt.Println("База данных успешно создана!")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", login, pass, host, port, dbName),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no change")

			return
		}

		panic(err)
	}

	fmt.Println("migrations successfully migrated")
}
