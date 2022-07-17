package Restaurants

type RestaurantsList struct { //структура всех ресторанов
	Id   int    `json:"id"`
	Name string `json:"name"` //название ресторана
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
	Id       int    `json:"id"`
	TableId  int    `json:"tableId"`
	UserId   int    `json:"userId"`
	TableNum int    `json:"tableNum"` //количество забронированных столов одного типа в данном ресторане
	Time     string `json:"time"`     // время бронирования столика
}

/*type BookingDate struct { // время бронирования столиков
	Id            int `json:"id"`
	TablesUsersId int `json:"tablesUsersId"` // для связи брони и времени брони
	//я пока плохо знаю как описать время, поэтому будет так пока что
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
	Hour  int `json:"hour"` //час начала брони
}
*/
