package handler

import (
	"financial-tracker-be/income"
	"net/http"

	"github.com/gin-gonic/gin"
)

type incomeHandler struct {
	income income.Service
}

func IncomeHandler(incomeService income.Service) *incomeHandler {
	return &incomeHandler{incomeService}
}

func (h *incomeHandler) GetAllIncome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "All Data Loaded",
	})
}

func (h *incomeHandler) GetIncomeByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "Get Income Data",
	})
}
func (h *incomeHandler) CreateIncome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "Income Created",
	})
}
func (h *incomeHandler) DeleteIncome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "Income Deleted",
	})
}
func (h *incomeHandler) UpdateIncome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "Income Update",
	})
}
