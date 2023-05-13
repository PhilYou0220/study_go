package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type Todo struct {
	gorm.Model
	Title  string `form:"title" json:"title" gorm:"column:title;comment:名称"` // 待办事项名称
	Status bool   `json:"status" gorm:"comment:0未完成1已完成"`             // 是否完成的状态
	Uid int64 `gorm:"comment:用户uid"`

}
type Account struct {
	gorm.Model
	Uid int64 `gorm:"column:uid;unique"` //用户唯一标识
	Name string `gorm:"type:varchar(255);column:name;not null;unique;comment:用户名称"` //用户名
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
	r.POST("/login",loginHandler)  //这些小写开头的函数能被访问到是因为是在同一个包（main）中
	r.POST("/register", registerHandler)
	g := r.Group("/api/v1",middleware)
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






