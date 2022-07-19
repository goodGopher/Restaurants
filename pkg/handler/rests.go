package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goodGopher/Restaurants"
)

func (h *Handler) createNewRest(c *gin.Context) {
	var input Restaurants.RestaurantsList
	if err := c.BindJSON(&input); err != nil {
		log.Fatalf("error rest parameters: %s", err.Error())
		return
	}

	id, err := h.service.CreateRestService(input)
	if err != nil {
		log.Fatalf("error CreateRestService: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllFreeRests(c *gin.Context) {

}

func (h *Handler) getRestById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id")) //получаем id из строки запроса
	if err != nil {
		log.Fatalf("error rest id: %s", err.Error())
		return
	}

	restById, err := h.service.GetRestByIdService(id)
	if err != nil {
		log.Fatalf("error GetRestByIdService: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, getOneRestResponse{
		Data: restById,
	})

}

func (h *Handler) updateRestById(c *gin.Context) {

}

func (h *Handler) deleteRestById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) //получаем id из строки запроса
	if err != nil {
		log.Fatalf("error rest id: %s", err.Error())
		return
	}

	err = h.service.DelRestByIdService(id)
	if err != nil {
		log.Fatalf("error DelRestByIdService: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, Status{
		Status: "ok",
	})

}

func (h *Handler) getAllRests(c *gin.Context) {

	allRests, err := h.service.GetAllRestsService()
	if err != nil {
		log.Fatalf("error GetAllRestsService: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, getRestsResponse{
		Data: allRests,
	})
}

func (h *Handler) getSortedByTimeRests(c *gin.Context) {

	allRestsSortedByTime, err := h.service.GetSortedByTimeRestsService()

	if err != nil {
		log.Fatalf("error getSortedByTimeRestsService: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, getRestsResponse{
		Data: allRestsSortedByTime,
	})
}

func (h *Handler) getSortedByCheckRests(c *gin.Context) {
	allRestsSortedByCheck, err := h.service.GetSortedByCheckRestsService()

	if err != nil {
		log.Fatalf("error GetSortedByCheckRestsService: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, getRestsResponse{
		Data: allRestsSortedByCheck,
	})

}
