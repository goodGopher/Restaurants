package main

import (
	"log"

	"github.com/goodGopher/Restaurants"
	"github.com/goodGopher/Restaurants/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(Restaurants.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}
