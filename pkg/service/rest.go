package service

import (
	"github.com/goodGopher/Restaurants"
	"github.com/goodGopher/Restaurants/pkg/repository"
)

type RestService struct {
	repo repository.Restaurant
}

func NewRestService(repo repository.Restaurant) *RestService {
	return &RestService{repo: repo}
}

func (r *RestService) CreateRestService(restaurant Restaurants.RestaurantsList) (int, error) {
	return r.repo.CreateRestRepo(restaurant)
}
