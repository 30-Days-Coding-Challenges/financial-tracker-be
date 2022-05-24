package main

import (
	"financial-tracker-be/handler"
	"financial-tracker-be/income"
	"financial-tracker-be/item"
	itemsource "financial-tracker-be/item_source"
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

	db.AutoMigrate(&income.Income{}, &itemsource.ItemSource{}, &item.Item{})

	itemSourceRepository := itemsource.ItemSourceRepository(db)
	itemSourceService := itemsource.ItemSourceService(itemSourceRepository)
	itemSourceHandler := handler.ItemSourceHandler(itemSourceService)

	incomeRepository := income.IncomeRepository(db)
	incomeService := income.IncomeService(incomeRepository)
	incomeHandler := handler.IncomeHandler(incomeService)

	itemRepository := item.ItemRepository(db)
	itemService := item.ItemService(itemRepository)
	itemHandler := handler.ItemHandler(itemService)

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":     "Muhammad Wage Juli Saputra",
			"position": "Software Engineer",
			"greet":    "Welcome to financial tracker service, enjoy your journey",
		})
	})

	v1.GET("/incomes", incomeHandler.GetAllIncome)
	v1.GET("/income/:id", incomeHandler.GetIncomeByID)
	v1.POST("/income", incomeHandler.CreateIncome)
	v1.DELETE("/income/:id", incomeHandler.DeleteIncome)
	v1.PUT("/income/:id", incomeHandler.UpdateIncome)

	v1.GET("/items", itemHandler.GetAllItem)
	v1.GET("/item/:id", itemHandler.GetItemByID)
	v1.POST("/item", itemHandler.CreateItem)
	v1.DELETE("/item/:id", itemHandler.DeleteItem)
	v1.PUT("/item/:id", itemHandler.UpdateItem)

	v1.POST("/item-source", itemSourceHandler.CreateItemSource)
	v1.GET("/item-sources", itemSourceHandler.GetAllItemSource)

	router.Run()
}
