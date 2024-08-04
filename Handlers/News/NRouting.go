package News

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"newsService1/Database"
)

func GetNewsByCategoryID(c *gin.Context) {
	categoryID := c.Param("id")
	news, err := fetchNewsByCategoryID(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if news == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No news found for this category"})
		return
	}

	c.JSON(http.StatusOK, news)
}

func GetNewsByID(c *gin.Context) {
	newsID := c.Param("news_id")
	news, err := fetchNewsByID(newsID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if news == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "News not found"})
		return
	}

	c.JSON(http.StatusOK, news)
}

func fetchNewsByCategoryID(categoryID string) ([]News, error) {
	rows, err := Database.DB.Query("SELECT id, title, description, categoryid , NDate ,full_description FROM news WHERE categoryid = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var news []News
	for rows.Next() {
		var item News
		if err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.CategoryID, &item.NDate, &item.FullDescription); err != nil {
			return nil, err
		}
		news = append(news, item)
	}

	return news, nil
}

func fetchNewsByID(newsID string) (*News, error) {
	row := Database.DB.QueryRow("SELECT id, title, description, categoryid, NDate ,full_description FROM news WHERE id = $1", newsID)
	var item News
	if err := row.Scan(&item.ID, &item.Title, &item.Description, &item.CategoryID, &item.NDate, &item.FullDescription); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}
