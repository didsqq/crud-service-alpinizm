package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/didsqq/crud-service-alpinizm/internal/services"
	"github.com/didsqq/crud-service-alpinizm/internal/storage"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := storage.NewMsSqlDB(storage.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		Encrypt:  viper.GetString("db.encrypt"),
		Password: viper.GetString("db.password"),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return
	}

	repo := storage.NewRepository(db)

	service := services.NewService(repo)

	log.Println("starting application")

	// application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	// go application.GRPCSrv.MustRun()

	//Graceful shutdown

	log.Println("Все заебись")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	// application.GRPCSrv.Stop()

	log.Printf("application stopped")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
