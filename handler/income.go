package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllIncome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "All Data Loaded",
	})
}

func GetIncomeByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "Get Income Data",
	})
}
func CreateIncome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "Income Created",
	})
}
func DeleteIncome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "Income Deleted",
	})
}
func UpdateIncome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "Income Update",
	})
}
