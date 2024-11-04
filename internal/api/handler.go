package api

import (
	"net/http"

	"github.com/DisVic/response-service/internal/repository"
	"github.com/DisVic/response-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(router *gin.Engine, aiService *service.AIService, repo *repository.Repository) {
	router.GET("/process", func(c *gin.Context) {
		query := c.Query("query")
		logrus.Infof("Запрос с параметром query: %s", query)

		if query == "" {
			logrus.Warn("Параметр query пуст.")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Параметр query не может быть пустым"})
			return
		}

		data, err := repo.GetDataByQuery(query)
		if err != nil {
			logrus.Errorf("Ошибка при поиске в БД: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка поиска данных"})
			return
		}

		if data != nil {
			logrus.Info("Данные найдены в БД, отправляем на фронтенд.")
			c.JSON(http.StatusOK, data)
			return
		}

		logrus.Info("Данные не найдены, обращаемся к AI-сервису.")
		response, err := aiService.FetchData(query)
		if err != nil {
			logrus.Errorf("Ошибка запроса к AI-сервису: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при запросе к AI-сервису"})
			return
		}
		logrus.Info("Ответ от AI-сервиса получен.")

		normalizedData := normalizeResponse(response)

		err = repo.SaveData(normalizedData)
		if err != nil {
			logrus.Errorf("Ошибка при сохранении данных в БД: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении данных"})
			return
		}

		logrus.Info("Данные отправлены на фронтенд.")
		c.JSON(http.StatusOK, normalizedData)
	})
}

func normalizeResponse(response *service.AIResponse) map[string]interface{} {
	if response == nil {
		logrus.Error("Ответ AI-сервиса пустой.")
		return nil
	}
	return map[string]interface{}{
		"descr": response.Descr,
		"text":  response.Text,
	}
}
