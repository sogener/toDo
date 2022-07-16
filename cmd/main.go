package main

import (
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"toDo"
	"toDo/package/handler"
	"toDo/package/repository"
	"toDo/package/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error while inititalization config %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("DATABASE_HOST"),
		Port:     viper.GetString("DATABASE_PORT"),
		Username: viper.GetString("DATABASE_USERNAME"),
		Password: viper.GetString("DATABASE_PASSWORD"),
		DBName:   viper.GetString("DATABASE_NAME"),
		SSLMode:  viper.GetString("DATABASE_SSL_MODE"),
	})

	if err != nil {
		log.Fatalf("error while connecting to the database %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(toDo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	return viper.ReadInConfig()
}
