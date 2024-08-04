package category

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"newsService1/Database"
	"strconv"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/categories", GetCategories)
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
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, category)
}

func fetchCategories() ([]Category, error) {
	rows, err := Database.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	categories := []Category{}
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
func fetchCategoryByID(id int) (Category, error) {
	var category Category
	row := Database.DB.QueryRow("SELECT id, name FROM categories WHERE id = $1", id)
	err := row.Scan(&category.ID, &category.Name)
	if err != nil {
		return category, err
	}
	return category, nil
}
