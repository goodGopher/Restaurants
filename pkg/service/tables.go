package service

import (
	"github.com/goodGopher/Restaurants"
	"github.com/goodGopher/Restaurants/pkg/repository"
)

type TableService struct {
	table repository.Tables
}

func NewTableService(table repository.Tables) *TableService {
	return &TableService{table: table}
}

func (t *TableService) CreateNewTableService(table Restaurants.TableList, rest_id int) (int, error) {
	return t.table.CreateNewTableRepo(table, rest_id)
}

func (t *TableService) GetFreeTablesService(rest_id int) ([]Restaurants.TableList, error) {
	return t.table.GetFreeTablesRepo(rest_id)
}

func (t *TableService) GetFreePlacesService(rest_id int) (int, int, int, error) {
	return t.table.GetFreePlacesRepo(rest_id)
}
