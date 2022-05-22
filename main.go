package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Muhammad Wage Juli Saputra",
			"position": "Software Engineer",
			"greet": "Welcome to this service, enjoy your journey",
		})
	})

	router.Run()
}