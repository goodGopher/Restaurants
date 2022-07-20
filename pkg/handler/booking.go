package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goodGopher/Restaurants"
)

func (h *Handler) createNewBooking(c *gin.Context) {
	var input Restaurants.BookingRequest
	if err := c.BindJSON(&input); err != nil {
		log.Fatalf("error booking parameters: %s", err.Error())
		return
	}

	id, err := h.service.CreateBookingService(input)
	if err != nil {
		log.Fatalf("error CreateBookingService: %s", err.Error())
		return
	}
	if id != nil {
		c.JSON(http.StatusOK, BookingResponse{
			ids: id,
		})
	} else {
		c.JSON(http.StatusOK, Status{
			Status: "no",
		})
	}
}

func (h *Handler) getBookingById(c *gin.Context) {

}

func (h *Handler) updateBookingById(c *gin.Context) {

}

func (h *Handler) deleteBookingById(c *gin.Context) {

}
