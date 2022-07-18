package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goodGopher/Restaurants"
)

func (h *Handler) createNewRest(c *gin.Context) {
	var input Restaurants.RestaurantsList

	if err := c.BindJSON(&input); err != nil {
		log.Fatalf("error rest parameters: %s", err.Error()) // нормальные логи завести
		return
	}

	id, err := h.service.CreateRestService(input)
	if err != nil {
		log.Fatalf("error CreateRestService: %s", err.Error()) // нормальные логи завести
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllFreeRests(c *gin.Context) {

}

func (h *Handler) getRestById(c *gin.Context) {

}

func (h *Handler) updateRestById(c *gin.Context) {

}

func (h *Handler) deleteRestById(c *gin.Context) {

}

func (h *Handler) getAllRests(c *gin.Context) {

}

func (h *Handler) getSortedByTimeRests(c *gin.Context) {

}

func (h *Handler) getSortedByCheckRests(c *gin.Context) {

}
