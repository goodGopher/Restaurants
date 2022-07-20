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
		booking := router.Group("/booking")
		{
			booking.POST("/", h.createNewBooking) // ?? создание брони в ресторане
			//booking.GET("/")                            //получение информации о всех бронированиях в ресторане
		//	booking.GET("/:id", h.getBookingById)       //получение информации о конкретном бронировании в ресторане
		//	booking.PUT("/:id", h.updateBookingById)    //обновление информации о конкретном бронировании в ресторане
		//	booking.DELETE("/:id", h.deleteBookingById) //удаление информации о конкретном бронировании в ресторане
		}

		restsActions := api.Group("/resAct")
		{
			restsActions.POST("/", h.createNewRest)       // ++ создание нового ресторана
		//	restsActions.GET("/", h.getAllFreeRests)      // получение списка всех ресторанов в которых есть свободные места с учетом условий
			restsActions.GET("/:id", h.getRestById)       // ++ получение имени конкретного ресторана по id
			restsActions.PUT("/:id", h.updateRestById)    // обновление имени конкретного ресторана по id
			restsActions.DELETE("/:id", h.deleteRestById) // ?? удалить столы ++ удаление конкретного ресторана по id

			tablesActions := restsActions.Group(":id/tabAct")
			{
				tablesActions.POST("/", h.createNewTable) // ++ создание нового типа столов в ресторане
				//	tablesActions.GET("/:peop_num", h.getTablesTypeByNum)    // получение информации о конкретном типе столов в ресторане
			//	tablesActions.PUT("/:tab_id", h.updateTablesTypeById) // обновление информации о конкретном типе столов в ресторане
				//	tablesActions.DELETE("/:peop_num", h.delTablesTypeByNum) // удаление конкретного типа столов в ресторане
			}
		}

		restsInfo := api.Group("/resInf")
		{
			restsInfo.GET("/all", h.getAllRests)                     // ++ получение списка всех ресторанов
			restsInfo.GET("/sortedByTime", h.getSortedByTimeRests)   // ++ сортировка подходящих ресторанов по ср.времени ожидания
			restsInfo.GET("/sortedByCheck", h.getSortedByCheckRests) // ++ сортировка подходящих ресторанов по ср.чеку

			tablesInfo := restsInfo.Group(":id/tabInf")
			{
				tablesInfo.GET("/allTables", h.getAllTablesInRest) // ++ вывод количества столов
				tablesInfo.GET("/allPlaces", h.getAllPlacesInRest) // ++ вывод количества  свободных мест в ресторане
			}

		}

	}
	return router
}
