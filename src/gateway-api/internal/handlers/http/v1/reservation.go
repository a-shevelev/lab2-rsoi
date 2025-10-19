package handlers

import (
	"gateway-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReservationHandler struct {
	Service *service.ReservationService
}

func NewReservationHandler(service *service.ReservationService) *ReservationHandler {
	return &ReservationHandler{Service: service}
}

func (h *ReservationHandler) RegisterRoutes(rg *gin.RouterGroup) {
	routes := rg.Group("/reservations")
	routes.GET("/", h.GetReservations)
}

func (h *ReservationHandler) GetReservations(c *gin.Context) {
	username := c.GetHeader("X-User-Name")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-User-Name header required"})
		return
	}

	reservations, err := h.Service.Get(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservations)
}
