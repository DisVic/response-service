package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var db *sql.DB
var log = logrus.New()
var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

func main() {
	// Загрузить переменные окружения
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	// Подключение к базе данных
	initDB()

	// Инициализация роутера
	router := gin.Default()

	// Маршруты
	router.POST("/data", saveData)

	// Запуск сервера
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}

func initDB() {
	var err error
	db, err = sql.Open("postgres", getDBConnectionString())
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	log.Info("Подключение к базе данных успешно")
	// Проверка соединения
	if err = db.Ping(); err != nil {
		log.Fatalf("База данных недоступна: %v", err)
	}
	log.Info("Подключение к базе данных успешно")
}

func getDBConnectionString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

type RequestData struct {
	Description string `json:"description" binding:"required"`
	Textik      string `json:"textik" binding:"required"`
}

func saveData(c *gin.Context) {
	var data RequestData
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Error("Ошибка валидации JSON:", err)
		c.JSON(400, gin.H{"error": "Неверный формат данных"})
		return
	}

	log.Info("Получен запрос для сохранения данных:", data)

	query := psql.Insert("data").Columns("description", "textik").Values(data.Description, data.Textik).Suffix("RETURNING id")
	var id int
	err := query.RunWith(db).QueryRow().Scan(&id)
	if err != nil {
		log.Error("Ошибка сохранения данных в БД:", err)
		c.JSON(500, gin.H{"error": "Ошибка сервера"})
		return
	}

	log.Info("Данные успешно сохранены с id:", id)
	c.JSON(200, gin.H{"id": id})
}

func init() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
}
