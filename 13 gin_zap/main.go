package main

import (
	"gin_zap/logger"
	"gin_zap/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


/*https://www.liwenzhou.com/posts/Go/zap-in-gin/
目的让zap记录gin本
*/

func main()  {
	//初始化日志模块
	err:=logger.Init()
	if err!=nil{
		panic(err)
	}


	gin.SetMode(gin.ReleaseMode) //设置为生产模式 不会显示一些debug日志
	//r := gin.Default() // 默认是debug模式
	r := gin.New() //自己写一个new函数
	//注册两个自定义的中间件
	r.Use(middleware.GinLogger(zap.L()),middleware.GinRecovery(zap.L(),false))
	r.GET("/hello", func(c *gin.Context) {
		zap.L().Info("req hello")
		panic("123") //有提示但是不会报错
	})

	r.Run(":8001")

}