package main

import (
	"financial-tracker-be/handler"
	"financial-tracker-be/income"
	"financial-tracker-be/item"
	itemsource "financial-tracker-be/item_source"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
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

	itemRepository := item.ItemRepository(db)
	itemService := item.ItemService(itemRepository)
	itemHandler := handler.ItemHandler(itemService)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	v1 := router.Group("/v1")

	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":     "Muhammad Wage Juli Saputra",
			"position": "Software Engineer",
			"greet":    "Welcome to financial tracker service, enjoy your journey",
		})
	})

	v1.GET("/items", itemHandler.GetAllItem)
	v1.GET("/item/:id", itemHandler.GetItemByID)
	v1.POST("/item", itemHandler.CreateItem)
	v1.DELETE("/item/:id", itemHandler.DeleteItem)
	v1.PUT("/item/:id", itemHandler.UpdateItem)

	v1.POST("/item-source", itemSourceHandler.CreateItemSource)
	v1.GET("/item-sources", itemSourceHandler.GetAllItemSource)

	router.Run()
}
