package main

import (
	"log"

	"github.com/polyk005/todo-app"
	"github.com/polyk005/todo-app/pkg/handler"
	"github.com/polyk005/todo-app/pkg/repository"
	"github.com/polyk005/todo-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8008"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
