package main

import (
	"log"

	gotodo "github.com/grancc/go-to-do-app"
	"github.com/grancc/go-to-do-app/pkg/handler"
	"github.com/grancc/go-to-do-app/pkg/repository"
	"github.com/grancc/go-to-do-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initialization config: %s", err.Error())
	}
	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)
	srv := new(gotodo.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
