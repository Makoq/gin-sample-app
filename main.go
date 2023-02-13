package main

import (
	database "demo/database"
	routers "demo/routers"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	// mysql  初始化
	database.InitDB()
}
func main() {
	
	//路由初始化
	r := routers.InitRoutes()

	// 设置日志文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// 使用日志中间件
	r.Use(gin.Logger())
	// 设置静态文件夹
	r.Static("/static", "./static")
	//启动端口
	r.Run(":8082")
}
