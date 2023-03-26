package main

import "github.com/gin-gonic/gin"

//和django比较像 都是要返回render 一个页面

func main(){
	//返回一个默认的路由引擎
	r := gin.Default()

	//注册templates下所有的html文件
	r.LoadHTMLGlob("templates/*")

	//返回字符串
	r.GET("/test01", func(context *gin.Context) {
		context.String(200,"hello world")
	})

	//返回html code码 html名称 obj 修改静态文件不用重启程序
	r.GET("/test02", func(context *gin.Context) {
		context.HTML(200,"index.html",nil)
		
	})
	//返回json数据
	r.POST("/test03", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"user_id": 1001,
			"username": "you",
		})

	})
	//返回xml数据
	r.POST("/test04", func(context *gin.Context) {
		context.XML(200,gin.H{
			"user_id": 1001,
			"username": "you",
			"friends": []string{"a","b","c"},
		})

	})

	//返回protobuf数据（后面补）

	//返回静态文件（配置）

	//设置静态资源的路径 参数一路由 参数二 本地真实路径
	r.Static("/static","./statics")


	r.Run("127.0.0.1:8081")

}
