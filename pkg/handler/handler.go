package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/goodGopher/Restaurants/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{

		rests := api.Group("/rests")
		{
			rests.POST("/", h.createNewRest)       // ++ создание нового ресторана
			rests.GET("/", h.getAllFreeRests)      // получение списка всех ресторанов в которых есть свободные места с учетом условий
			rests.GET("/:id", h.getRestById)       // ++ получение имени конкретного ресторана по id
			rests.PUT("/:id", h.updateRestById)    // обновление имени конкретного ресторана по id
			rests.DELETE("/:id", h.deleteRestById) // удаление конкретного ресторана по id

			rests.GET("/all", h.getAllRests)                     // ++ получение списка всех ресторанов
			rests.GET("/sortedByTime", h.getSortedByTimeRests)   // ++ сортировка подходящих ресторанов по ср.времени ожидания
			rests.GET("/sortedByCheck", h.getSortedByCheckRests) // ++ сортировка подходящих ресторанов по ср.чеку

			/*tables := rests.Group(":id/tables")
			{
				tables.POST("/")      //создание нового стола в ресторане
				tables.GET("/")       // вывод количества столов и свободных мест в ресторане
				tables.GET("/:id")    // получение информации о конкретном типе столов в ресторане
				tables.PUT("/:id")    // обновление информации о конкретном типе столов в ресторане
				tables.DELETE("/:id") // удаление конкретного типа столов в ресторане
			}
			*/
			booking := rests.Group(":id/booking")
			{
				booking.POST("/", h.createNewBooking) //создание брони в ресторане
				//booking.GET("/")                            //получение информации о всех бронированиях в ресторане
				booking.GET("/:id", h.getBookingById)       //получение информации о конкретном бронировании в ресторане
				booking.PUT("/:id", h.updateBookingById)    //обновление информации о конкретном бронировании в ресторане
				booking.DELETE("/:id", h.deleteBookingById) //удаление информации о конкретном бронировании в ресторане
			}
		}
	}
	return router
}
