package main

import "github.com/gin-gonic/gin"
import . "route/route"
//服务端对客户端发起各种请求方式的处理


//func book(ctx *gin.Context){
//
//	ctx.HTML(200,"login.html",nil)
//}


func main()	{
	//获取引擎对象，即路由对象
	r:=gin.Default()

	//加载模板文件
	//r.LoadHTMLGlob("templates/*")



	//初始化路由
	InitBookRoute(r)
	//启动：默认本机8080端口
	r.Run("127.0.0.1:8081")
}

