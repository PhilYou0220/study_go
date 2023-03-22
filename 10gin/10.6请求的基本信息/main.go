package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)


func main() {
	r:=gin.Default()
	//路由映射函数同一个路由，不同的方法执行不同的逻辑
	r.Any("/book", func(context *gin.Context) {

		//获取请求方式
		fmt.Println("method:",context.Request.Method)
		fmt.Println("method:",context.Request.URL)
		fmt.Println(context.FullPath())

		context.JSON(200, gin.H{
			"msg": "查询成功1",
		})

	})



	r.Run("127.0.0.1:8081")

}
