package service

import (
	"github.com/goodGopher/Restaurants"
	"github.com/goodGopher/Restaurants/pkg/repository"
)

type Restaurant interface {
	CreateRestService(input Restaurants.RestaurantsList) (int, error)
	GetAllRestsService() ([]Restaurants.RestaurantsList, error)
	GetSortedByTimeRestsService() ([]Restaurants.RestaurantsList, error)
	GetSortedByCheckRestsService() ([]Restaurants.RestaurantsList, error)
	GetRestByIdService(id int) (Restaurants.RestaurantsList, error)
	DelRestByIdService(id int) error
}

type Tables interface {
	CreateNewTableService(table Restaurants.TableList, rest_id int) (int, error)
	GetFreeTablesService(rest_id int) ([]Restaurants.TableList, error)
	GetFreePlacesService(rest_id int) (int, int, int, error)
}

type Booking interface {
	CreateBookingService(bookingReq Restaurants.BookingRequest) ([]int, error)
}

type Service struct {
	Restaurant
	Tables
	Booking
}

func NewService(repos *repository.Repos) *Service {
	return &Service{
		Restaurant: NewRestService(repos),
		Tables:     NewTableService(repos),
	}
}
