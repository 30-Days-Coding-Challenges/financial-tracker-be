package handler

import (
	itemsource "financial-tracker-be/item_source"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type itemSourceHandler struct {
	itemSource itemsource.Service
}

func ItemSourceHandler(itemSourceService itemsource.Service) *itemSourceHandler {
	return &itemSourceHandler{itemSourceService}
}

func (h *itemSourceHandler) CreateItemSource(c *gin.Context) {
	var itemSourceRequest itemsource.ItemSourceRequest

	err := c.ShouldBindJSON(&itemSourceRequest)
	source, err := h.itemSource.CreateItemSource(itemSourceRequest)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Source item Created",
		"data":    source,
	})
}

func (h *itemSourceHandler) GetAllItemSource(c *gin.Context) {

	allitemSource, _ := h.itemSource.GetAllItemSource()
	c.JSON(http.StatusOK, gin.H{
		"data": allitemSource,
	})
}

func (h *itemSourceHandler) DeleteItemSource(c *gin.Context) {

	itemSourceID := c.Param("id")

	itemSource, err := h.itemSource.GetItemSourceByID(itemSourceID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Item source not found",
			"error": err,
		})
	}

	deletedItemSource, err := h.itemSource.DeleteItemSource(itemSource)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message" : "Internal server error",
			"error": err,
		})
	}
	
	c.JSON(http.StatusOK, gin.H{
		"status" : "OK",
		"message": "Item source Deleted",
		"data" : deletedItemSource,
	})
}
