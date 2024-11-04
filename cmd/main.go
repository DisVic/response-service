package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/DisVic/response-service/config"
	"github.com/DisVic/response-service/internal/api"
	"github.com/DisVic/response-service/internal/repository"
	"github.com/DisVic/response-service/internal/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // импортируем pq для работы с PostgreSQL
)

func main() {
	cfg, err := config.LoadConfig() // Загружает конфигурацию (настройки базы данных и URL AI-сервиса)
	if err != nil {
		log.Fatalf("Не удалось загрузить конфигурацию: %v", err)
	}

	// Подключение к базе данных с использованием pq
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer db.Close()

	// Проверка соединения
	if err := db.PingContext(context.Background()); err != nil {
		log.Fatalf("Не удалось проверить соединение с базой данных: %v", err)
	}

	router := gin.Default()
	aiService := service.NewAIService(cfg.AIServiceURL)
	repo := repository.NewRepository(db)
	api.SetupRoutes(router, aiService, repo) // Настройка маршрутов

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
