package repository

import (
	"github.com/goodGopher/Restaurants"
	"github.com/jmoiron/sqlx"
)

type Restaurant interface {
	CreateRestRepo(restaurant Restaurants.RestaurantsList) (int, error)
	GetAllRestsRepo() ([]Restaurants.RestaurantsList, error)
	GetSortedByTimeRestsRepo() ([]Restaurants.RestaurantsList, error)
	GetSortedByCheckRestsRepo() ([]Restaurants.RestaurantsList, error)
	GetRestByIdRepo(id int) (Restaurants.RestaurantsList, error)
}

type Booking interface {
}

type Repos struct {
	Restaurant
	Booking
}

func NewRepository(db *sqlx.DB) *Repos {
	return &Repos{
		Restaurant: NewRestRepo(db),
	}
}
