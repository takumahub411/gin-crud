package controllers

import (
	"gin-curd/initializers"
	"gin-curd/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddMovie(c *gin.Context) {
	var body models.Movie

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "読み込み失敗",
		})
		return
	}

	result := initializers.DB.Create(&body)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "投稿失敗",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "投稿完了",
	})
}

func GetMovies(c *gin.Context) {
	var movies []models.Movie

	result := initializers.DB.Find(&movies)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "取得失敗",
		})
		//c.AbortWithError(http.StatusNotFound, result.Error)
	}
	c.JSON(http.StatusOK, &movies)

}

func GetMovie(c *gin.Context) {
	var movies models.Movie

	id := c.Param("id")

	result := initializers.DB.Find(&movies, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "取得失敗",
		})
		//c.AbortWithError(http.StatusNotFound, result.Error)
	}
	c.JSON(http.StatusOK, &movies)

}

func UpdateMovie(c *gin.Context) {
	var movies models.Movie

	// update request
	var updateInfo models.RequestMovie

	id := c.Param("id")

	if err := c.ShouldBindJSON(&updateInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "読み込み失敗",
		})
		return
	}

	if result := initializers.DB.First(&movies, id); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "取得失敗",
		})
		return
	}

	movies.Title = updateInfo.Title
	movies.LeadingRole = updateInfo.LeadingRole
	movies.Description = updateInfo.Description
	movies.Stars = updateInfo.Stars

	// update
	initializers.DB.Save(&movies)

	c.JSON(http.StatusOK, &movies)
}

func DeleteMovie(c *gin.Context) {
	var movies models.Movie

	id := c.Param("id")

	result := initializers.DB.First(&movies, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "取得失敗",
		})
		//c.AbortWithError(http.StatusNotFound, result.Error)
	}
	initializers.DB.Delete(&movies)

	c.JSON(http.StatusOK, &movies)
}
