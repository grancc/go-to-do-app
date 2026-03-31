package main

// @title           To Do App API
// @version         1.0
// @description     API for todo lists. Protected routes: Authorization header with two words, e.g. Bearer plus JWT.
// @host            localhost:8080
// @BasePath        /
// @schemes         http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	gotodo "github.com/grancc/go-to-do-app"
	"github.com/grancc/go-to-do-app/pkg/handler"
	"github.com/grancc/go-to-do-app/pkg/repository"
	"github.com/grancc/go-to-do-app/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialization config: %s", err.Error())
	}

	db, err := repository.NewPOstgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initilizqation db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)
	srv := new(gotodo.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("ToDoApp start")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logrus.Print("ToDoApp shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on close db connection: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
