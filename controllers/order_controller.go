package controllers

import (
	"net/http"

	"restaurant-api/models"
	"restaurant-api/repositories"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderRepo *repositories.OrderRepo
}

func NewOrderController(or *repositories.OrderRepo) *OrderController {
	return &OrderController{OrderRepo: or}
}

type createOrderReq struct {
	Items []models.OrderItem `json:"items" binding:"required"`
}

func (oc *OrderController) Create(c *gin.Context) {
	var in createOrderReq
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}
	userID := uid.(int)

	order := &models.Order{
		UserID: userID,
		Status: "pending",
		Total:  0,
	}

	if err := oc.OrderRepo.Create(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	total := 0.0
	for i := range in.Items {
		it := &in.Items[i]
		it.OrderID = order.ID

		if err := oc.OrderRepo.AddItem(it); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		total += float64(it.Quantity) * it.Price
	}

	_, _ = oc.OrderRepo.DB.Exec(
		`UPDATE orders SET total=$1 WHERE id=$2`,
		total,
		order.ID,
	)

	c.JSON(http.StatusCreated, gin.H{
		"order_id": order.ID,
		"total":    total,
	})
}
