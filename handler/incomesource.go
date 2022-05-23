package handler

import (
	incomesource "financial-tracker-be/income_source"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type incomeSourceHandler struct {
	incomeSource incomesource.Service
}

func IncomeSourceHandler(incomeSourceService incomesource.Service) *incomeSourceHandler {
	return &incomeSourceHandler{incomeSourceService}
}

func (h *incomeSourceHandler) CreateIncomeSource(c *gin.Context) {
	var incomeSourceRequest incomesource.IncomeSourceRequest

	err := c.ShouldBindJSON(&incomeSourceRequest)
	source, err := h.incomeSource.CreateIncomeSource(incomeSourceRequest)

	fmt.Println(incomeSourceRequest)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Source Income Created",
		"data":    source,
	})
}

func (h *incomeSourceHandler) GetAllIncomeSource(c *gin.Context) {

	allIncomeSource, _ := h.incomeSource.GetAllIncomeSource()
	c.JSON(http.StatusOK, gin.H{
		"data": allIncomeSource,
	})
}
