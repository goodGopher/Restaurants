package main

import (
	"log"
	"os"

	"github.com/goodGopher/Restaurants"
	"github.com/goodGopher/Restaurants/pkg/handler"
	"github.com/goodGopher/Restaurants/pkg/repository"
	"github.com/goodGopher/Restaurants/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %v", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
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
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
