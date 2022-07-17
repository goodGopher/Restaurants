package repository

type Restaurant interface {
}

type Booking interface {
}

type Repos struct {
	Restaurant
	Booking
}

func NewRepository() *Repos {
	return &Repos{}
}
