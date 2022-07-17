package service

import "github.com/goodGopher/Restaurants/pkg/repository"

type Restaurant interface {
}

type Booking interface {
}

type Service struct {
	Restaurant
	Booking
}

func NewService(repos *repository.Repos) *Service {
	return &Service{}
}
