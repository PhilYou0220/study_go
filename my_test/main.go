package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

//type RestrictedArea struct {
//	ID            uint
//	AreaID        int
//	Name          string
//	TargetType    bool `gorm:"default:0"`
//	AreaPoints    string
//	CenterPoints  string
//	DateType      bool      `gorm:"default:0"`
//	Deleted       bool      `gorm:"default:0"`
//	PassType      bool      `gorm:"default:0"`
//	ConfigSpeed   float64   `gorm:"default:0.0"`
//	CreatedBy     int       `gorm:"default:0"`
//	CreatedTime   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
//	ConditionType int8      `gorm:"default:0"`
//	HasWhiteCar   int8      `gorm:"default:0"`
//	StopLimit     int       `gorm:"default:0"`
//}

func DBInit() *gorm.DB {
	// 创建数据库链接
	//parseTime=True：将 MySQL 中的时间类型（如 datetime、timestamp 等）解析为 Go 中的时间类型（time.Time）。如果不设置此参数，查询出来的时间类型会是 []uint8，需要手动转换为 time.Time。
	//loc=Local：指定时区为本地时区。如果不设置此参数，时间会被转换为 UTC 时间，需要手动转换为本地时区
	//DSN（Data Source Name），用于指定数据库的连接信息。
	dsn := "epuser:epuser@123-Test@tcp(106.75.138.97:3306)/ep?charset=utf8mb4&parseTime=True&loc=Local"
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
	//db.AutoMigrate(&RestrictedArea{})

	return db

}

var db = DBInit()


func main() {
	db.First()

	fmt.Println(db)


}
