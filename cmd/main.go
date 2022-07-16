package main

import (
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

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(toDo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	return viper.ReadInConfig()
}
