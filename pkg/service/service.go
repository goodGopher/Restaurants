package service

import (
	"github.com/goodGopher/Restaurants"
	"github.com/goodGopher/Restaurants/pkg/repository"
)

type Restaurant interface {
	CreateRestService(input Restaurants.RestaurantsList) (int, error)
}

type Booking interface {
}

type Service struct {
	Restaurant
	Booking
}

func NewService(repos *repository.Repos) *Service {
	return &Service{
		Restaurant: NewRestService(repos),
	}
}
