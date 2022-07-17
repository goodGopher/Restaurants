package main

import (
	"log"

	"github.com/goodGopher/Restaurants"
	"github.com/goodGopher/Restaurants/pkg/handler"
	"github.com/goodGopher/Restaurants/pkg/repository"
	"github.com/goodGopher/Restaurants/pkg/service"
)

func main() {

	/*для передачи запроса по слоям обработки
	выстроим зависимости между обрабатывающими интерфейсами*/
	repo := repository.NewRepository()  // работа с БД
	service := service.NewService(repo) // основная логика
	handler := handler.NewHandler(service)

	srv := new(Restaurants.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}
