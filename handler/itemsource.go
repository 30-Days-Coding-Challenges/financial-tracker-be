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
