package main

import (
	"log"

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

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
