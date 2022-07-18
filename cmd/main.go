package main

import (
	"log"

	"github.com/goodGopher/Restaurants"
	"github.com/goodGopher/Restaurants/pkg/handler"
	"github.com/goodGopher/Restaurants/pkg/repository"
	"github.com/goodGopher/Restaurants/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %v", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5437",
		Username: "postgres",
		Password: "zxcvbn",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	/*для передачи запроса по слоям обработки
	выстроим зависимости между обрабатывающими интерфейсами*/
	repo := repository.NewRepository(db) // работа с БД
	service := service.NewService(repo)  // основная логика
	handler := handler.NewHandler(service)

	srv := new(Restaurants.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
