package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"newsService1/Database"
	"newsService1/Handlers/News"
	"newsService1/Handlers/category"
)

func main() {
	Database.DbInit()
	router := gin.Default()

	router.GET("/categories", category.GetCategories)
	router.GET("/categories/:id", category.GetCategoryByID)
	router.GET("/categories/:id/news", News.GetNewsByCategoryID)
	router.GET("/categories/:id/news/:news_id", News.GetNewsByID)

	router.Run(":8080")
}
