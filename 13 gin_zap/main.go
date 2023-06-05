package main

import "github.com/gin-gonic/gin"

func main()  {
	gin.SetMode(gin.ReleaseMode) //设置为生产模式
	r := gin.Default() // 默认是debug模式

	r.Run(":8001")

}