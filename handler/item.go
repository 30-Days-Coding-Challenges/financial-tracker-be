package handler

import (
	"financial-tracker-be/item"
	"financial-tracker-be/utils"
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
	itemID := c.Param("id")

	if utils.IsStringUUID(itemID) != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": "Invalid Format ID",
		})
		return
	}

	item, err := h.item.GetItemByID(itemID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "ERROR",
			"message": "Item not found",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "OK",
		"data":    item,
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
	itemID := c.Param("id")

	if utils.IsStringUUID(itemID) != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": "Invalid Format ID",
		})
		return
	}

	_, err := h.item.GetItemByID(itemID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "ERROR",
			"message": "Item not found",
			"error":   err,
		})
		return
	}

	errDelete := h.item.DeleteItem(itemID)

	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "INTERNAL SERVER ERROR",
			"message": errDelete,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Item Deleted",
		"data":    "ID: " + itemID,
	})
}

func (h *itemHandler) UpdateItem(c *gin.Context) {
	itemID := c.Param("id")

	if utils.IsStringUUID(itemID) != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": "Invalid Format ID",
		})
		return
	}

	_, err := h.item.GetItemByID(itemID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "ERROR",
			"message": "Item not found",
			"error":   err,
		})
		return
	}

	var itemReq item.ItemRequest

	c.ShouldBindJSON(&itemReq)
	updatedItem, err := h.item.UpdateItem(itemID, itemReq)

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Item Updated",
		"data":    updatedItem,
	})
}
