package main

//下载
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

//创建模型类，抽象公共字段
type BaseModel struct {
	ID int `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(32);unique;not null"`
	CreateTime *time.Time `gorm:"autoCreateTime"`
	UpdateTime *time.Time `gorm:"autoCreateTime"`
}

type Teacher struct{
	BaseModel
	Sno int
	Pwd string `gorm:"type:varchar(100);not null"`
	Tel string `gorm:"type:char(11);"`
	Birth *time.Time
	Remark string `gorm:"type:varchar(255);"`
}
//默认引擎表名会加复数，若不想使用默认引擎  函数名称为TableName 继承Teacher结构体 返回 什么 表的名字就是什么
//func (t Teacher)TableName() string  {
//	return "teacher"
//}
type Class struct {
	BaseModel //继承基础表
	Tno int
	Pwd string `gorm:"type:varchar(100);not null"`
	Birth *time.Time
	Remark string `gorm:"type；varchar(255);"`

}
func main() {
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger}) //配置了日志器
	fmt.Println(db, err)

	// 关闭数据库链接
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	//迁移模型类 模型转换为sql 创建表执行创表语句
	db.AutoMigrate(&Teacher{})

}

