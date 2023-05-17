package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

//3业务处理每个接口都这么写
//3.1获取参数
//3.2编写业务逻辑
//3.3返回结果
//新增
func createTodoHandler(ctx *gin.Context) {
	//3.1获取参数
	var todo Todo //待会把参数放在todo里
	err := ctx.ShouldBind(&todo)
	if err != nil {
		fmt.Println("invalid param", err)
		ctx.JSON(200, gin.H{
			"code": -1,
			"msg":  "无效的参数",
		})
		return //没有返回值
	}

	//3.2编写业务逻辑
	// 获取ctx内的用户身份
	v,_ := ctx.Get(CtxUIdKey) //返回的是一个接口类型
	uid := v.(int64)
	fmt.Println("uid是：",uid)
	if uid<=0{
		ctx.JSON(200, gin.H{
			"code": 1,
			"msg":  "登录异常",
		})
		return
	}
	//拿到uid之后赋值给结构体
	todo.Uid=uid
	res := db.Create(&todo)
	if res.Error != nil {
		fmt.Println("db.Create failed", err)
		ctx.JSON(200, gin.H{
			"code": -1,
			"msg":  "服务端异常", // 正常来说最好不要直接把后端错误返回给前端
		})
		return
	}
	//3.3返回结果
	ctx.JSON(200,gin.H{
		"code":1,
		"msg":"新增成功",
	})
}

//查询
func getTodoHandler(ctx *gin.Context) {
	//3.1获取参数 查询全部没有参数获取

	//3.2 处理业务逻辑 全部应该是个切片
	var todo []Todo
	//查询用户
	v,_ := ctx.Get(CtxUIdKey) //返回的是一个接口类型
	uid := v.(int64)
	if uid<=0{
		ctx.JSON(200, gin.H{
			"code": 1,
			"msg":  "token异常",
		})
		return
	}

	if err:=db.Model(&Todo{}).Where("uid=?",uid).Find(&todo).Error;err!=nil{ //注意是返回值的error
		fmt.Println("查询出错！！！",err)
		ctx.JSON(200, gin.H{
			"code": -1,
			"msg":  "服务端异常，请稍后再试",
		})
		return
	}
	//3.3返回结果
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": todo,
	})
}

//修改
func updateTodoHandler(ctx *gin.Context) {
	//1获取参数
	var todo Todo
	if err := ctx.ShouldBind(&todo); err != nil {
		fmt.Println("db.Create failed", err)
		ctx.JSON(200, gin.H{
			"code": -1,
			"msg":  "服务端异常", // 正常来说最好不要直接把后端错误返回给前端
		})
		return
	}

	//2业务逻辑
	// 先检查数据（todo.ID）是否存在
	// 拿请求传过来的id去数据库里查询是否存在这条记录
	//查询用户
	v,_ := ctx.Get(CtxUIdKey) //返回的是一个接口类型
	uid := v.(int64)
	if uid<=0{
		ctx.JSON(200, gin.H{
			"code": 1,
			"msg":  "token异常",
		})
		return
	}

	if err := db.Where("uid=?",uid).First(&Todo{}, todo.ID).Error; err != nil {
		// if err == gorm.ErrRecordNotFound {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没有这条记录
			ctx.JSON(200, gin.H{
				"code": 1,
				"msg":  "无效的参数",
			})
			return
		}
		// 其他错误
		ctx.JSON(200, gin.H{
			"code": -1,
			"msg":  "服务端异常，请稍后再试",
		})
		return
	}

	// 代码能执行到这里，说明数据库中确实存在 todo.ID 对应的记录
	// 接下来就去更新这条记录
	res:=db.Model(&Todo{}).Where("id = ? and uid=?", todo.ID,uid).Update("status", todo.Status)
	if res.Error != nil{
		fmt.Println("updateTodoHandler：更新至数据库失败", res.Error)
		ctx.JSON(200, gin.H{
			"code": 1,
			"msg":  "服务端异常，请稍后再试",
		})
		return
	}
	// 3. 返回响应
	ctx.JSON(200, gin.H{
		"code": 1,
		"msg":  "success",
	})

}

func deleteTodoHandler(ctx *gin.Context) {
	//1获取参数
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	var todo Todo
	if err:=ctx.ShouldBind(&todo);err!=nil{
		fmt.Println("参数错误！！！")
		ctx.JSON(200,gin.H{
			"code": -1,
			"msg":  "参数异常",

		})
	}
	//2业务处理

	v,_ := ctx.Get(CtxUIdKey) //返回的是一个接口类型
	uid := v.(int64)
	if uid<=0{
		ctx.JSON(200, gin.H{
			"code": 1,
			"msg":  "token异常",
		})
		return
	}

	if err:=db.Model(&Todo{}).Where("uid=?",uid).First(&todo).Error;err!=nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没有这条记录
			ctx.JSON(200, gin.H{
				"code": 1,
				"msg":  "没有这条记录",
			})
			return
		}
		// 其他错误
		ctx.JSON(200, gin.H{
			"code": 1,
			"msg":  "服务端异常，请稍后再试",
		})
		return
	}
	// 代码能执行到这里，说明数据库中确实存在 todo.ID 对应的记录
	// 接下来就去更新这条记录
	if err := db.Where("uid=?",uid).Delete(&Todo{}, id).Error; err != nil {
		ctx.JSON(200, gin.H{
			"code": 1,
			"msg":  "服务端异常，请稍后再试",
		})
		return
	}
	// 3. 返回响应
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
	})
}