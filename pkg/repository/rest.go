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
	query := fmt.Sprintf("INSERT INTO %s (rest_name) values ($1) RETURNING id", restaurantsList)
	row := r.db.QueryRow(query, restaurant.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
