package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{

		rests := api.Group("/rests")
		{
			rests.POST("/", h.createNewRest)       // создание нового ресторана
			rests.GET("/", h.getAllFreeRests)      // получение списка всех ресторанов в которых есть свободные места с учетом условий
			rests.GET("/:id", h.getRestById)       // получение имени конкретного ресторана по id
			rests.PUT("/:id", h.updateRestById)    // обновление имени конкретного ресторана по id
			rests.DELETE("/:id", h.deleteRestById) // удаление конкретного ресторана по id

			rests.GET("/all", h.getAllRests)                    // получение списка всех ресторанов
			rests.GET("/sortedByTime", h.getSortedByTimeRests)  // сортировка подходящих ресторанов по ср.времени ожидания
			rests.GET("/sortedByBill", h.getSortedByCheckRests) // сортировка подходящих ресторанов по ср.чеку

			booking := rests.Group(":id/booking")
			{
				booking.POST("/", h.createNewBooking) //создание брони в ресторане
				//	booking.GET( "/") //получение информации о всех бронированиях в ресторане
				booking.GET("/:id", h.getBookingById)       //получение информации о конкретном бронировании в ресторане
				booking.PUT("/:id", h.updateBookingById)    //обновление информации о конкретном бронировании в ресторане
				booking.DELETE("/:id", h.deleteBookingById) //удаление информации о конкретном бронировании в ресторане
			}

			/*
				tables := rests.Group( ":id/tables"){
					tables.POST( "/") //создание нового стола в ресторане
					tables.GET( "/") // вывод количества столов и свободных мест в ресторане
					tables.GET( "/:id") // получение информации о конкретном типе столов в ресторане
					tables.PUT( "/:id") // обновление информации о конкретном типе столов в ресторане
					tables.DELETE( "/:id") // удаление конкретного типа столов в ресторане


					bookingTables := tables.Group( ":id/bookingTables"){
						bookingTables.POST( "/") // создание брони на определенный тип столов
					//	bookingTables.GET( "/") // получение информации о брони на определенный тип столов
						bookingTables.PUT( "/") // обновление брони на определенный тип столов
						bookingTables.DELETE( "/") // удаление брони на определенный тип столов

					}


				}*/
		}
	}
	return router
}
