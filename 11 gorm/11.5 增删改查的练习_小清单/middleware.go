package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
)

func middleware(ctx *gin.Context) {
	//1获取参数
	auth_param := ctx.Request.Header.Get("Authorization")
	//ctx.GetHeader("Authorization")
	fmt.Println(auth_param)
	if auth_param == "" {
		ctx.JSON(http.StatusOK, Resp{
			Code: 401,
			Msg:  "无请求头，权限错误",
		})
		ctx.Abort() // 终止请求 首先，使用ctx.Abort()可以确保在停止请求的执行之前，可以执行一些必要的清理操作，例如释放资源、关闭连接等。而使用return可能会导致这些清理操作无法执行，从而导致资源泄漏或其他问题。
		//其次，使用ctx.Abort()可以确保在停止请求的执行之前，可以设置HTTP响应的状态码和响应头，从而更好地控制响应的内容。而使用return则无法直接控制响应的状态码和响应头，需要通过其他方式来实现。
		return
	}

	//ctx.Next() //next写个函数名称执行这个函数，执行完了再回到主流程来
	//2业务逻辑
	part := strings.Split(auth_param," ")
	fmt.Println(part[0])
	fmt.Println(reflect.TypeOf(part[0]))
	if part[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized,Resp{
			Code: 401,
			Msg:  "非合法的请求头",
		})
		ctx.Abort()
		return
	}
	mc,err :=ParseToken(part[1])
	if err != nil{
		ctx.JSON(http.StatusUnauthorized,Resp{
			Code: 401,
			Msg:  "无效的token",
		})
		ctx.Abort()
		return
	}
	//把用户名和uid写入ctx对象 k,v
	ctx.Set(CtxUIdKey,mc.Uid)
	ctx.Set(CtxNameKey,mc.Name)


	//3返回参数
}
