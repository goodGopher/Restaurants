package Restaurants

type RestaurantsList struct { //структура всех ресторанов
	Id      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"rest_name" binding:"required"`    //название ресторана
	AvTime  string `json:"av_time" db:"av_time" binding:"required"`   // среднее время ожидания
	AvCheck int    `json:"av_check" db:"av_check" binding:"required"` // средний чек
}

type RestaurantsTables struct { // структура для связи столиков и ресторанов
	Id           int `json:"id"`
	RestaurantId int `json:"restaurant_id" db:"restaurant_id" `
	TableId      int `json:"table_id" db:"table_id"`
}

type TableList struct { //структура всех столиков в ресторане
	Id         int `json:"id" db:"id"`
	TableNum   int `json:"table_num" db:"table_num"`     // количество столов в ресторане
	MaxPersons int `json:"max_persons" db:"max_persons"` // максимальное число людей за столом
	FreeTables int `json:"free_tables" db:"free_tables"` // количество свободных столов
}

type BookingList struct { //структура для ,бронирования (связи столиков и пользователей)
	Id          int    `json:"id" db:"id"`
	TableId     int    `json:"table_id" db:"table_id"`
	UserId      int    `json:"user_id" db:"user_id"`
	TableNum    int    `json:"table_num" db:"table_num"` //количество забронированных столов одного типа в данном ресторане
	BookingDate string `json:"booking_date" db:"booking_date"`
	BookingTime string `json:"booking_time" db:"booking_time"`
}

type BookingRequest struct { // структура для получения запроса на бронирование
	RestName    string `json:"rest_name" db:"rest_name"`
	PeopNum     int    `json:"peop_num" db:"peop_num"`
	UserName    string `json:"user_name" db:"user_name"`
	UserPhone   string `json:"user_phone" db:"user_phone"`
	BookingDate string `json:"booking_date" db:"booking_date"`
	BookingTime string `json:"booking_time" db:"booking_time"`
}
