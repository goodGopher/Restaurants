package Restaurants

type User struct {
	Id    int    `json:"-"`
	Name  string `json:"name"`
	Phone string `jsin:"phone"`
}
