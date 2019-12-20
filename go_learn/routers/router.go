package routers

import (
	"go_learn/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadRouters(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Status": 0,
				"data":   "Hello World!",
			})
	})
	//登录
	router.POST("/login", controllers.Login)
	// 查询发文列表
	router.GET("/article", controllers.GetPosts)

}
