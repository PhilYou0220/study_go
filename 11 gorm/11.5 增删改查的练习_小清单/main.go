package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
)

type Todo struct {
	gorm.Model
	Title  string `form:"title" json:"title" gorm:"column:title;comment:名称"` // 待办事项名称
	Status bool   `json:"status" gorm:"comment:0未完成1已完成"`             // 是否完成的状态
}
type Account struct {
	gorm.Model

	Name string `gorm:"type:varchar(255);column:name;not null;unique;comment:用户名称"` //唯一标识
	NickName string `gorm:"type:varchar(255);comment:用户昵称"` //昵称
	PassWord string `gorm:"column:pass_word;comment:密码;type:text"`
	Status *bool `gorm:"column:status;comment:状态"`

}
var db *gorm.DB

func DbInit()  {
	dsn := "root:youfei123@tcp(1.15.60.231:3306)/css?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), //标准输出流（os.Stdout） 在每条日志之间添加了一个回车换行符（"\r\n"），时间戳格式为标准格式（log.LstdFlags）
		logger.Config{
			LogLevel: logger.Info, //记录的日志级别
		},
	)
	var err error                                                         //先定义个error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger}) //使用了全局变量所以不使用:=
	if err != nil {
		//return db //全局变量不需要返回
		fmt.Println("连接数据库失败！！！")
		panic(err)
	} else {
		fmt.Println("连接数据库成功！！！")

	}

}

func main() {
	//1获取路由对象并进行 路由匹配响应函数
	r := gin.Default()

	//1.1  路由匹配响应函数
	r.LoadHTMLFiles("./index.html")
	r.Static("/static", "./static")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	g := r.Group("/api/v1")
	{
		g.POST("/todo", createTodoHandler)
		g.PUT("/todo", updateTodoHandler)
		g.GET("/todo", getTodoHandler)
		g.DELETE("/todo/:id", deleteTodoHandler) // 路由参数
	}
	//2初始化数据库并获取数据库对象
	DbInit()
	//db.AutoMigrate(&Todo{})
	//db.AutoMigrate(&Account{})

	//3业务处理每个接口都这么写
	//3.1获取参数
	//3.2编写业务逻辑
	//3.3返回结果

	//4 启动gin服务
	if err := r.Run(":8001"); err != nil {
		fmt.Println("启动服务失败！")
	}
}





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
	if err:=db.Model(&Todo{}).Find(&todo).Error;err!=nil{ //注意是返回值的error
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
	if err := db.First(&Todo{}, todo.ID).Error; err != nil {
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
	res:=db.Model(&Todo{}).Where("id = ?", todo.ID).Update("status", todo.Status)
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
	if err:=db.Model(&Todo{}).First(&todo).Error;err!=nil{
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
	if err := db.Delete(&Todo{}, id).Error; err != nil {
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
