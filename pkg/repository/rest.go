package repository

import (
	"fmt"

	"github.com/goodGopher/Restaurants"
	"github.com/jmoiron/sqlx"
)

type RestRepo struct {
	db *sqlx.DB
}

func NewRestRepo(db *sqlx.DB) *RestRepo {
	return &RestRepo{db: db}
}

func (r *RestRepo) CreateRestRepo(restaurant Restaurants.RestaurantsList) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (rest_name, av_time, av_check) values ($1, $2, $3) RETURNING id", restaurantsList)
	row := r.db.QueryRow(query, restaurant.Name, restaurant.AvTime, restaurant.AvCheck)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *RestRepo) GetAllRestsRepo() ([]Restaurants.RestaurantsList, error) {
	var allRests []Restaurants.RestaurantsList
	query := fmt.Sprintf("SELECT * FROM %s ", restaurantsList)
	err := r.db.Select(&allRests, query)

	return allRests, err
}

func (r *RestRepo) GetSortedByTimeRestsRepo() ([]Restaurants.RestaurantsList, error) {
	var allRestsSortedByTime []Restaurants.RestaurantsList
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY av_time", restaurantsList)
	err := r.db.Select(&allRestsSortedByTime, query)

	return allRestsSortedByTime, err
}

func (r *RestRepo) GetSortedByCheckRestsRepo() ([]Restaurants.RestaurantsList, error) {
	var allRestsSortedByCheck []Restaurants.RestaurantsList
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY av_check", restaurantsList)
	err := r.db.Select(&allRestsSortedByCheck, query)

	return allRestsSortedByCheck, err
}

func (r *RestRepo) GetRestByIdRepo(id int) (Restaurants.RestaurantsList, error) {
	var restById Restaurants.RestaurantsList
	query := fmt.Sprintf("SELECT * FROM %s WHERE id= $1", restaurantsList)
	err := r.db.Get(&restById, query, &id)

	return restById, err
}
