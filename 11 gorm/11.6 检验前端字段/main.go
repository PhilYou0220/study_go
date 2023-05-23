package main

import "github.com/gin-gonic/gin"

type Login struct {
	Username string `json:"username,omitempty" form:"username" binding:"required" ` //binding:"required" 必传项
	Password string `json:"password" form:"password" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"` //binding:"email"` 符合email规则
	Age      int    `json:"age" form:"age" binding:"required,gte=1,lte=100"` // 大于等于1 小于等于100
}

func loginfunc(c *gin.Context) {
	//1 获取参数
	var login Login
	if err:=c.ShouldBind(&login);err!=nil{
		c.JSON(400,gin.H{
			"code": 401,
			"msg": err.Error(),//
		})
		return
	}
	//2 业务逻辑
}


func main(){
	//初始化gin对象
	r:=gin.Default()
	r.POST("/login",loginfunc)
	r.Run(":9999")
}

