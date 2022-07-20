package handler

import (
	"level0/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type Handler struct {
	storage service.StorageData
}

func NewHandler(storage service.StorageData) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) GetOrder(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 8)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	order := h.storage.Get(uint(id))

	c.JSON(http.StatusOK, order)
}

func (h *Handler) GetAllOrders(c *gin.Context) {
	orders, err := h.storage.MEM.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, orders)
}
