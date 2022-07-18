package repository

import (
	"github.com/goodGopher/Restaurants"
	"github.com/jmoiron/sqlx"
)

type Restaurant interface {
	CreateRestRepo(restaurant Restaurants.RestaurantsList) (int, error)
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
