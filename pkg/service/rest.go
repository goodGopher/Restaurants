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

func (r *RestService) GetAllRestsService() ([]Restaurants.RestaurantsList, error) {
	return r.repo.GetAllRestsRepo()
}

func (r *RestService) GetSortedByTimeRestsService() ([]Restaurants.RestaurantsList, error) {
	return r.repo.GetSortedByTimeRestsRepo()
}

func (r *RestService) GetSortedByCheckRestsService() ([]Restaurants.RestaurantsList, error) {
	return r.repo.GetSortedByCheckRestsRepo()
}

func (r *RestService) GetRestByIdService(id int) (Restaurants.RestaurantsList, error) {
	return r.repo.GetRestByIdRepo(id)
}

func (r *RestService) DelRestByIdService(id int) error {
	return r.repo.DelRestByIdRepo(id)
}
