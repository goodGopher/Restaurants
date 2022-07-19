package Restaurants

type RestaurantsList struct { //структура всех ресторанов
	Id      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"rest_name" binding:"required"`    //название ресторана
	AvTime  string `json:"av_time" db:"av_time" binding:"required"`   // среднее время ожидания
	AvCheck int    `json:"av_check" db:"av_check" binding:"required"` // средний чек
}

type RestaurantsTables struct { // структура для связи столиков и ресторанов
	Id           int `json:"id"`
	RestaurantId int `json:"restaurantId"`
	TableId      int `json:"tableId"`
}

type TableList struct { //структура всех столиков в ресторане
	Id         int `json:"id"`
	TableNum   int `json:"tableNum"`    // количество столов в ресторане
	MaxPersons int `json:"max_persons"` // максимальное число людей за столом
	FreeTables int `json:"free_tables"` // количество свободных столов
}

type BookingList struct { //структура для ,бронирования (связи столиков и пользователей)
	Id          int    `json:"id"`
	TableId     int    `json:"table_id"`
	UserId      int    `json:"user_id"`
	TableNum    int    `json:"table_num"` //количество забронированных столов одного типа в данном ресторане
	BookingDate string `json:"booking_date"`
	BookingTime string `json:"booking_time"`
}
