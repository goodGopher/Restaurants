package Restaurants

type UserList struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"name" binding:"required"`
	Phone string `jsin:"phone" binding:"required"`
}
