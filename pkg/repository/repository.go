package repository

import "github.com/jmoiron/sqlx"

type Restaurant interface {
}

type Booking interface {
}

type Repos struct {
	Restaurant
	Booking
}

func NewRepository(db *sqlx.DB) *Repos {
	return &Repos{}
}
