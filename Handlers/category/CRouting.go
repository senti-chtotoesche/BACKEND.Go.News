package category

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"newsService1/Database"
	"strconv"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/categories", GetCategories)
	router.GET("/categories/:id", GetCategoryByID)
}

func GetCategories(c *gin.Context) {
	categories, err := fetchCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func GetCategoryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	category, err := fetchCategoryByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, category)
}

func fetchCategories() ([]Category, error) {
	var categories []Category
	if err := Database.DB.Find(&categories).Error; err != nil {
		log.Println("Error fetching categories:", err)
		return nil, err
	}
	return categories, nil
}

func fetchCategoryByID(id int) (Category, error) {
	var category Category
	if err := Database.DB.First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return category, nil
		}
		log.Println("Error fetching category by ID:", err)
		return category, err
	}
	return category, nil
}
