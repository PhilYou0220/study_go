//main.go
package main

import "github.com/gin-gonic/gin"
func getuser(ctx *gin.Context){
	ctx.JSON(200,gin.H{
		"username":"you",
	})
}
func main()	{
	//获取引擎对象，即路由对象
	r:=gin.Default()

	//路由映射函数
	r.GET("/user",getuser)
	//启动：默认本机8080端口 类似django的runserver
	r.Run("127.0.0.1:8081")
}
