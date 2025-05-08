package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/didsqq/crud-service-alpinizm/internal/handler"
	"github.com/didsqq/crud-service-alpinizm/internal/repository"
	"github.com/didsqq/crud-service-alpinizm/internal/service"
)

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	db, err := repository.NewMsSqlDB(repository.Config{
		Host:     getEnvVar("DB_HOST"),
		Port:     getEnvVar("DB_PORT"),
		Username: getEnvVar("DB_USERNAME"),
		DBName:   getEnvVar("DB_DATABASE"),
		Encrypt:  getEnvVar("DB_ENCRYPT"),
		Password: getEnvVar("DB_SA_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return
	}

	repo := repository.NewUnitOfWork(db)

	service := service.NewService(repo)

	handler := handler.NewHandler(service)

	r := handler.InitRoutes()

	srv := &http.Server{
		Addr:    ":" + getEnvVar("APP_PORT"),
		Handler: r,
	}

	log.Println("starting application")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	log.Println("application started")

	//Graceful shutdown

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("failed to stop server: %v", err)

		return
	}

	log.Printf("application stopped")
}

func getEnvVar(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Ошибка: переменная окружения %s не найдена", key)
	}
	return value
}
