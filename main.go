package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"newsService1/Database"
	"newsService1/Handlers/News"
	"newsService1/Handlers/category"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	Database.DbInit(dsn)

	router := gin.Default()

	router.GET("/categories", category.GetCategories)
	router.GET("/categories/:id", category.GetCategoryByID)
	router.GET("/categories/:id/news", News.GetNewsByCategoryID)
	router.GET("/categories/:id/news/:news_id", News.GetNewsByID)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
