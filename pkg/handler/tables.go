package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goodGopher/Restaurants"
)

func (h *Handler) createNewTable(c *gin.Context) {
	var table Restaurants.TableList
	var rest_id int

	if err := c.BindJSON(&table); err != nil {
		log.Fatalf("error table parameters: %s", err.Error())
		return
	}

	rest_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalf("error Rest id: %s", err.Error())
		return
	}

	table_id, err := h.service.CreateNewTableService(table, rest_id)
	if err != nil {
		log.Fatalf("error CreateNewTableService: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": table_id,
	})
}

func (h *Handler) getAllTablesInRest(c *gin.Context) {
	var tables []Restaurants.TableList
	var rest_id int

	rest_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalf("error Rest id: %s", err.Error())
		return
	}

	tables, err = h.service.GetFreeTablesService(rest_id)
	if err != nil {
		log.Fatalf("error GetFreeTablesService: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, FreeTables{
		FreeTables: tables,
	})
}

func (h *Handler) getAllPlacesInRest(c *gin.Context) { //выводит суммарное максимальное количество свободных столов, всех столов и свободных мест
	var SumTableNum1 int
	var SumMaxPersons1 int
	var SumFreeTables1 int
	var rest_id int

	rest_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalf("error Rest id: %s", err.Error())
		return
	}

	SumFreeTables1, SumTableNum1, SumMaxPersons1, err = h.service.GetFreePlacesService(rest_id)
	if err != nil {
		log.Fatalf("error GetFreeTablesService: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, FreePlaces{
		SumFreeTables: SumFreeTables1,
		SumTableNum:   SumTableNum1,
		SumMaxPersons: SumMaxPersons1,
	})
}

func (h *Handler) getTablesTypeById(c *gin.Context) {

}

func (h *Handler) updateTablesTypeById(c *gin.Context) {

}

func (h *Handler) delTablesTypeById(c *gin.Context) {

}
