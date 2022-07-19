package repository

import (
	"fmt"

	"github.com/goodGopher/Restaurants"
	"github.com/jmoiron/sqlx"
)

type TableRepo struct {
	db *sqlx.DB
}

func NewTableRepo(db *sqlx.DB) *TableRepo {
	return &TableRepo{db: db}
}

func (t *TableRepo) CreateNewTableRepo(table Restaurants.TableList, rest_id int) (int, error) {

	tq, err := t.db.Begin() // начало серии неразрывных действий с таблицами
	if err != nil {
		return 0, err
	}

	var table_id int

	//создаем стол в tableList
	createTable := fmt.Sprintf("INSERT INTO %s (table_num, max_persons, free_tables) values ($1, $2, $3) RETURNING id", tableList)
	row := tq.QueryRow(createTable, table.TableNum, table.MaxPersons, table.FreeTables)
	if err := row.Scan(&table_id); err != nil {
		tq.Rollback() //откат действий с таблицами
		return 0, err
	}

	//связываем рестораны с одним типом столов
	bindTableAndRest := fmt.Sprintf("INSERT INTO %s (restaurant_id, table_id) values ($1, $2) ", restaurantsTables)
	_, err = tq.Exec(bindTableAndRest, rest_id, table_id)
	if err != nil {
		tq.Rollback() //откат действий с таблицами
		return 0, err
	}

	return table_id, tq.Commit()
}

func (t *TableRepo) GetFreeTablesRepo(rest_id int) ([]Restaurants.TableList, error) {

	var allFreeTables []Restaurants.TableList

	result := fmt.Sprintf("SELECT a1.id, a1.table_num, a1.max_persons, a1.free_tables FROM %s AS a1 INNER JOIN %s AS a2 ON  a1.id = a2.table_id WHERE a1.free_tables > 0 AND a2.restaurant_id = $1", tableList, restaurantsTables)
	err := t.db.Select(&allFreeTables, result, &rest_id)

	return allFreeTables, err

}

func (t *TableRepo) GetFreePlacesRepo(rest_id int) (int, int, int, error) {

	var SumTableNum int
	var SumMaxPersons int
	var SumFreeTables int

	result := fmt.Sprintf("SELECT SUM(a1.table_num) table_num, SUM(a1.max_persons) max_persons,SUM( a1.free_tables) free_tables FROM %s AS a1 INNER JOIN %s AS a2 ON  a1.id = a2.table_id WHERE a1.free_tables > 0 AND a2.restaurant_id = $1 ", tableList, restaurantsTables)
	row := t.db.QueryRow(result, rest_id)
	if err := row.Scan(&SumTableNum, &SumMaxPersons, &SumFreeTables); err != nil {
		return 0, 0, 0, err
	}

	return SumFreeTables, SumTableNum, SumMaxPersons, nil

}
