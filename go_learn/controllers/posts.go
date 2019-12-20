package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go_learn/models"
)

//GetPosts 获取所有的文章
func GetPosts(c *gin.Context) {
	posts := models.GetPosts()
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   posts,
	})
}