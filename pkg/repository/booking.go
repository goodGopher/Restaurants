package repository

import (
	"fmt"
	"log"

	"github.com/goodGopher/Restaurants"
	"github.com/jmoiron/sqlx"
)

type BookingRepo struct {
	db *sqlx.DB
}

func NewBookingRepo(db *sqlx.DB) *BookingRepo {
	return &BookingRepo{db: db}
}

func (nbk *BookingRepo) CheckTablesIdsRepo(RestName string, PeopNum int) (map[int]int, error) {
	//получение id ресторана по имени
	var IdRestByName int

	query := fmt.Sprintf("SELECT * FROM %s WHERE rest_name= $1 RETURNING id", restaurantsList)
	err := nbk.db.Get(&IdRestByName, query, &RestName)
	if err != nil {
		log.Fatalf("error get rest by name: %s", err.Error())
		return nil, err
	}

	//получение информации о столах в ресторане
	var tr Repos
	var RestTables []Restaurants.TableList //все столы
	var FreePlaces int                     //свободные места

	RestTables, err = tr.Tables.GetFreeTablesRepo(IdRestByName)
	if err != nil {
		log.Fatalf("error GetFreeTablesRepo: %s", err.Error())
		return nil, err
	}

	FreePlaces, _, _, err = tr.Tables.GetFreePlacesRepo(IdRestByName)
	if err != nil {
		log.Fatalf("error GetFreeTablesRepo: %s", err.Error())
		return nil, err
	}

	var resultIds map[int]int

	if FreePlaces < PeopNum {
		log.Fatalf("error CheckTablesIdsRepo(): not enough free places: %s", err.Error()) //эта ошибка должна пресекаться выдачей возможных вариантов
		return nil, err
	} else {
		//ищем столы с достаточным количеством мест
		needId := -1

		//вычисления без оптимизации
		for _, v := range RestTables {
			if (v.MaxPersons > PeopNum) && (v.FreeTables > 0) { //есть достаточно большой свободный стол
				needId = v.Id
			}
		}
		if needId >= 0 {
			return map[int]int{needId: 1}, nil
		} else {
			sum := 0 // суммарное число мест у сдвинутых столов
			for _, v := range RestTables {
				if sum < PeopNum {
					sum += v.FreeTables * v.MaxPersons
					resultIds[v.Id] = v.FreeTables
				} else {
					break
				}
			}
			return resultIds, nil
		}
	}
}

func (nbk *BookingRepo) CreateUserForBookingRepo(UserName string, UserPhone string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (users_name, phone) values ($1, $2) RETURNING id", userList)
	row := nbk.db.QueryRow(query, UserName, UserPhone)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (nbk *BookingRepo) CreateBookingRepo(BkList Restaurants.BookingList) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (table_id, user_id, table_num,booking_date,booking_time) values ($1, $2, $3, $4, $5) RETURNING id", bookingList)
	row := nbk.db.QueryRow(query, BkList.TableId, BkList.UserId, BkList.TableNum, BkList.BookingDate, BkList.BookingTime)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
