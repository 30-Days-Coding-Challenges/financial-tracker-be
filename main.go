package main

import (
	"financial-tracker-be/handler"
	"financial-tracker-be/income"
	incomesource "financial-tracker-be/income_source"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/financial_tracker?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Error", err)
	}

	fmt.Println("DB Connected")

	db.AutoMigrate(&income.Income{}, &incomesource.IncomeSource{})

	incomeSourceRepository := incomesource.SourceIncomeRepository(db)
	incomeSourceService := incomesource.IncomeSourceService(incomeSourceRepository)
	incomeSourceHandler := handler.IncomeSourceHandler(incomeSourceService)

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":     "Muhammad Wage Juli Saputra",
			"position": "Software Engineer",
			"greet":    "Welcome to financial tracker service, enjoy your journey",
		})
	})

	v1.GET("/incomes", handler.GetAllIncome)
	v1.GET("/income/:id", handler.GetIncomeByID)
	v1.POST("/income", handler.CreateIncome)
	v1.DELETE("/income/:id", handler.DeleteIncome)
	v1.PUT("/income/:id", handler.UpdateIncome)
	v1.POST("/income-source", incomeSourceHandler.CreateIncomeSource)
	v1.GET("/income-sources", incomeSourceHandler.GetAllIncomeSource)

	router.Run()
}
