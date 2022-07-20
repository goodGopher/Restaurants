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
	DelRestByIdRepo(id int) error
}

type Tables interface {
	CreateNewTableRepo(table Restaurants.TableList, rest_id int) (int, error)
	GetFreeTablesRepo(rest_id int) ([]Restaurants.TableList, error)
	GetFreePlacesRepo(rest_id int) (int, int, int, error)
}
type Booking interface {
	CheckTablesIdsRepo(RestName string, PeopNum int) (map[int]int, error)
	CreateUserForBookingRepo(UserName string, UserPhone string) (int, error)
	CreateBookingRepo(BkList Restaurants.BookingList) (int, error)
}

type Repos struct {
	Restaurant
	Tables
	Booking
}

func NewRepository(db *sqlx.DB) *Repos {
	return &Repos{
		Restaurant: NewRestRepo(db),
		Tables:     NewTableRepo(db),
		Booking:    NewBookingRepo(db),
	}
}
