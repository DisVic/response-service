package main

import (
	"github.com/DisVic/response-service/config"
	"github.com/DisVic/response-service/internal/api"
	"github.com/DisVic/response-service/internal/repository"
	"github.com/DisVic/response-service/internal/service"

	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	cfg, err := config.LoadConfig() // Загружает конфигурацию (настройки базы данных и URL AI-сервиса)
	if err != nil {
		log.Fatalf("Не удалось загрузить конфигурацию: %v", err)
	}

	dbPool, err := pgxpool.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer dbPool.Close()

	router := gin.Default()
	aiService := service.NewAIService(cfg.AIServiceURL)
	repo := repository.NewRepository(dbPool)
	api.SetupRoutes(router, aiService, repo) // Настройка маршрутов

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
