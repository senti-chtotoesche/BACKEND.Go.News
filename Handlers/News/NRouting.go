package News

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	var news []News
	if err := Database.DB.Where("categoryid = ?", categoryID).Find(&news).Error; err != nil {
		return nil, err
	}
	return news, nil
}

func fetchNewsByID(newsID string) (*News, error) {
	var item News
	if err := Database.DB.First(&item, "id = ?", newsID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}
