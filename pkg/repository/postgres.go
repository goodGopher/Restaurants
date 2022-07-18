package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	restaurantsList = "restaurants_list"

	userList = "user_list"

	tableList = "table_list"

	restaurantsTables = "restaurants_tables"

	bookingList = "booking_list"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) { //читаем параметры из файла
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	//проверка подключения к базе
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
