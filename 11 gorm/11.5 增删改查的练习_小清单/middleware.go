package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func middleware(ctx *gin.Context) {
	//1获取参数
	auth_param := ctx.Request.Header.Get("Authorization")
	//ctx.GetHeader("Authorization")
	fmt.Println(auth_param)
	//ctx.Next() //next写个函数名称执行这个函数，执行完了再回到主流程来
	//2业务逻辑
	//3返回参数
}
