package api

import (
	"net/http"

	"github.com/DisVic/response-service/internal/repository"
	"github.com/DisVic/response-service/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, aiService *service.AIService, repo *repository.Repository) {
	router.GET("/process", func(c *gin.Context) {
		response, err := aiService.FetchData()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных от AI"})
			return
		}

		normalizedData := normalizeResponse(response)
		err = repo.SaveData(normalizedData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сохранения данных"})
			return
		}

		c.JSON(http.StatusOK, normalizedData)
	})
}

func normalizeResponse(response *service.AIResponse) map[string]interface{} {
	return map[string]interface{}{
		"description": response.Description,
		"text":        response.Text,
	}
}
