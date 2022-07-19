package handler

import "github.com/goodGopher/Restaurants"

type Status struct {
	Status string `json:"status"`
}

//Rests responses
type getRestsResponse struct { //дополнительная структура для вывода списка ресторанов
	Data []Restaurants.RestaurantsList `json:"data"`
}
type getOneRestResponse struct { //дополнительная структура для вывода списка ресторанов
	Data Restaurants.RestaurantsList `json:"data"`
}

//Tables responses
type FreeTables struct {
	FreeTables []Restaurants.TableList `json:"free_tables"`
}

type FreePlaces struct {
	SumFreeTables int `json:"sum_free_tables"`
	SumTableNum   int `json:"sum_table_num"`
	SumMaxPersons int `json:"sum_max_persons"`
}
