package main

import (
	"github.com/gin-gonic/gin"
	"go_learn/database"
	"go_learn/routers"
)

func main() {
	r := gin.Default()

	// 连接数据库
	database.InitDB()

	// 加载路由
	routers.LoadRouters(r)

	// 静态资源
	r.Static("/static", "./static")

	r.Run(":8081") //http://localhost:8081/ping
	// {"data":"Hello World","status":0}
}
