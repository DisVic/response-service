package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Симулируем AI-сервис, который использует параметр query
	router.GET("/mock-ai-service", func(c *gin.Context) {
		query := c.Query("query")
		response := gin.H{
			"descr": "Описание для запроса: " + query,
			"text":  "Текст для запроса: " + query,
		}
		if query == "" {
			response = gin.H{
				"descr": "Описание (запрос пуст)",
				"text":  "Текст (запрос пуст)",
			}
		}
		c.JSON(http.StatusOK, response)
	})

	router.Run(":8081")
}
