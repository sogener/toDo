package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"toDo"
	"toDo/package/handler"
	"toDo/package/repository"
	"toDo/package/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error while inititalization config %s", err.Error())
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
		logrus.Fatalf("error while connecting to the database %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(toDo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	return viper.ReadInConfig()
}
