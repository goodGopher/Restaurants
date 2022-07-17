package Restaurants

type UserList struct {
	Id    int    `json:"-"`
	Name  string `json:"name"`
	Phone string `jsin:"phone"`
}
