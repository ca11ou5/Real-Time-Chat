package main

import (
	"github.com/spf13/viper"
	"log"
	"realtimeChat/internal/delivery/http"
	"realtimeChat/internal/repository"
	"realtimeChat/internal/repository/postgres"
	"realtimeChat/internal/server"
	"realtimeChat/internal/service"
)

func main() {
	//Config initialization
	err := initConfig()
	if err != nil {
		log.Fatal("failed to initialize config")
	}

	//Database, Repository, Service, Handler initialization
	db := postgres.NewDatabase()
	repo := repository.NewRepository(db)
	serv := service.NewService(repo)
	hand := http.NewHandler(serv)

	//HTTP Server
	srv := server.NewServer(viper.GetString("server.port"), hand.InitRoutes())
	srv.Run()
}

func initConfig() error {
	viper.SetConfigFile("configs/main.yaml")
	return viper.ReadInConfig()
}
