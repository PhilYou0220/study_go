package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

//创建模型类，抽象公共字段，对于此模型类还有一个 gorm.Model的公共字段
type BaseModel struct {
	//gorm.Model
	ID         int        `gorm:"primaryKey"`
	Name       string     `gorm:"type:varchar(32);unique;not null"`
	CreateTime *time.Time `gorm:"autoCreateTime"`
	UpdateTime *time.Time `gorm:"autoCreateTime"`
}

type Teacher struct {
	BaseModel
	Tno    int
	Pwd    string `gorm:"type:varchar(100);not null"`
	Tel    string `gorm:"type:char(11);"`
	Birth  *time.Time
	Remark string `gorm:"type:varchar(255);"`
}

//默认引擎表名会加复数，若不想使用默认引擎  函数名称为TableName 继承Teacher结构体 返回 什么 表的名字就是什么
//func (t Teacher)TableName() string  {
//	return "teacher"
//}
type Class struct {
	BaseModel
	Num     int
	TutorID int
	Tutor   Teacher `gorm:"foreignKey:Tno"` //这样写让人一目了然
}
type Course struct {
	BaseModel
	Credit    int // 学分
	Period    int // 周期
	// 一对多 这个是建立外键的
	TeacherID int
	Teacher   Teacher
}

type Student struct {
	BaseModel
	Sno    int
	Pwd    string `gorm:"type:varchar(100);not null"`
	Tel    string `gorm:"type:char(11);"`
	Gender byte   `gorm:"default:1"` //相当于unit8
	Birth  *time.Time
	Remark string `gorm:"type:varchar(255);"`
	// 建立一对多的关联关系
	ClassID int
	Class   Class
	// 多对多:通过创建第三张表
	Courses []Course `gorm:"many2many:student2course;constraint:OnDelete:CASCADE;"`
}
var db = DBInit()
func DBInit() *gorm.DB {
	// 创建数据库链接
	//parseTime=True：将 MySQL 中的时间类型（如 datetime、timestamp 等）解析为 Go 中的时间类型（time.Time）。如果不设置此参数，查询出来的时间类型会是 []uint8，需要手动转换为 time.Time。
	//loc=Local：指定时区为本地时区。如果不设置此参数，时间会被转换为 UTC 时间，需要手动转换为本地时区
	//DSN（Data Source Name），用于指定数据库的连接信息。
	dsn := "root:youfei123@tcp(1.15.60.231:3306)/css?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), //标准输出流（os.Stdout） 在每条日志之间添加了一个回车换行符（"\r\n"），时间戳格式为标准格式（log.LstdFlags）
		logger.Config{
			LogLevel: logger.Info, //记录的日志级别
		},
	)

	//db接收*gorm.DB 类型的对象 gorm.Config{} 是 gorm.Open 函数的第二个参数，用于指定 GORM 库的一些配置
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger}) //配置了日志器
	//fmt.Println(db, err)
	//迁移模型类 模型转换为sql 创建表执行创表语句

	// 自动迁移
	db.AutoMigrate(&Course{})
	db.AutoMigrate(&Teacher{})
	db.AutoMigrate(&Class{})
	db.AutoMigrate(&Student{})

	return db

}
func add(ctx *gin.Context) {
	//添加记录
	//重要 orm映射 结构体对象 结构体对象映射表记录 因此很多时候都要操作结构体对象

	//实例化结构体
	//t1 := Teacher{BaseModel: BaseModel{Name: "youfei1"}, Tno: 1001, Pwd: "yuan", Tel: "110"}
	//fmt.Println(t1)
	//fmt.Println(t1.Name)
	//fmt.Println(t1.ID) //结构体的默认值 0

	////（1）增--单表单条记录 传结构体对象
	//db.Create(&t1) //很重要，传了地址过去，可以返回当时插入的值 如果不加&仅仅是个值拷贝罢了 注意返回的时结构体对象 值时从数据库拿到的值，若数据报错了,ID=0,插入失败MYsql 自增id也会自己增加
	//if t1.ID!=0 {
	//	fmt.Println("插入成功",t1.ID)
	//
	//}else {
	//	fmt.Println("插入后:",t1.ID)
	//}
	//
	//ctx.JSON(200, gin.H{
	//	"msg": t1,
	//})

	//增--前端增
	////可以使用ctx获得前端的值 通过postman发送
	//name := ctx.PostForm("name")
	//tno,_ := strconv.Atoi(ctx.PostForm("tno"))
	//pwd :=ctx.PostForm("pwd")
	//tel :=ctx.PostForm("tel")
	////赋值给结构体
	//t := Teacher{BaseModel: BaseModel{Name: name}, Tno: tno, Pwd: pwd, Tel: tel}
	//db.Create(&t)
	//ctx.JSON(200, gin.H{
	//	"msg": t,
	//})

	//（2）单表创建多条记录 传结构体切片[]struct
	//t1 := Teacher{BaseModel: BaseModel{Name: "rain"}, Tno: 1001, Pwd: "123"}
	//t2 := Teacher{BaseModel: BaseModel{Name: "eric"}, Tno: 1002, Pwd: "123"}
	//t3 := Teacher{BaseModel: BaseModel{Name: "yuan"}, Tno: 1003, Pwd: "123"}
	//teachers := []Teacher{t1, t2, t3}
	//db.Create(&teachers)
	//ctx.JSON(200, gin.H{
	//	"msg": teachers, //返回了一个{ [ {},{} ] } 不要在意返回了啥子，你把他当成[]Teacher{t1, t2, t3}就行了
	//})


	class01 := Class{BaseModel: BaseModel{Name: "软件1班"}, TutorID: 1}
	class02 := Class{BaseModel: BaseModel{Name: "软件2班"}, TutorID: 2}
	class03 := Class{BaseModel: BaseModel{Name: "计算机科学与技术1班"}, TutorID: 3}
	class04 := Class{BaseModel: BaseModel{Name: "计算机科学与技术2班"}, TutorID: 3}
	classes := []Class{class01, class02, class03, class04}
	db.Create(&classes)
	ctx.JSON(200, gin.H{
		"msg": classes, //返回了一个{ [ {},{} ] } 不要在意返回了啥子，你把他当成[]Teacher{t1, t2, t3}就行了
	})

}



func main() {

	r := gin.Default()
	//DBInit()
	fmt.Println(r)
	//程序启动起来之后 做路由匹配 不会重复执行上面的代码
	r.POST("/add", add)
	r.Run("127.0.0.1:8081")
	// 关闭数据库链接
	//sqlDB, err := db.DB()
	//if err != nil {
	//	panic(err)
	//}
	//defer sqlDB.Close()

}
