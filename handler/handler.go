package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order-service/domain"
	"order-service/usecase"
	"strconv"
)

type Handler struct {
	service *usecase.Service
}

func NewHandler() *Handler {
	return &Handler{service: usecase.NewService()}
}

// GET /orders
func (h *Handler) GetAllOrders(c *gin.Context) {
	orders, err := h.service.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve orders: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// GET /orders/:id
func (h *Handler) GetOrderByID(c *gin.Context) {
	idStr := c.Param("id")
	orderID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid order ID: " + err.Error(),
		})
		return
	}

	order, err := h.service.GetOrderByID(uint(orderID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Order not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, order)
}

// POST /orders
func (h *Handler) CreateOrder(c *gin.Context) {
	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	created, err := h.service.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

// PATCH /orders/:id
func (h *Handler) UpdateOrder(c *gin.Context) {
	idStr := c.Param("id")
	orderID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID: " + err.Error()})
		return
	}

	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	order.OrdetID = uint(orderID)

	updated, err := h.service.UpdateOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

// DELETE /orders/:id
func (h *Handler) DeleteOrder(c *gin.Context) {
	idStr := c.Param("id")
	orderID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID: " + err.Error()})
		return
	}

	err = h.service.DeleteOrder(uint(orderID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}
