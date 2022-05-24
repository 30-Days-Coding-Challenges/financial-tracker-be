package handler

import (
	"financial-tracker-be/item"
	"net/http"

	"github.com/gin-gonic/gin"
)

type itemHandler struct {
	item item.Service
}

func ItemHandler(itemService item.Service) *itemHandler {
	return &itemHandler{itemService}
}

func (h *itemHandler) GetAllItem(c *gin.Context) {

	items, _ := h.item.GetAllItem()

	c.JSON(http.StatusOK, gin.H{
		"data": items,
	})
}

func (h *itemHandler) GetItemByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "Get Item Data",
	})
}
func (h *itemHandler) CreateItem(c *gin.Context) {
	var itemReq item.ItemRequest

	err := c.ShouldBindJSON(&itemReq)
	newItem, err := h.item.CreateItem(itemReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Item Created",
		"data":    newItem,
	})
}
func (h *itemHandler) DeleteItem(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "Item Deleted",
	})
}
func (h *itemHandler) UpdateItem(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "Item Update",
	})
}
