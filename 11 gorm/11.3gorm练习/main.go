package main

import (
	//"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Teacher struct {
	Id          int        `gorm:"primaryKey;autoIncrement;comment:主键id"` //所谓蛇形复数
	Tno         int        `gorm:"default:0"`
	Name        string     `gorm:"type:varchar(10);not null"`
	Pwd         string     `gorm:"type:varchar(100);not null"`
	Tel         string     `gorm:"type:char(11);column:my_name"`
	Birth       *time.Time //它的零值（默认值）将是time.Time{}，而不是 nil，因为 time.Time 是值类型，它的默认值是其零值。如果你想要在这个字段中存储 NULL 值，就需要使用 *time.Time 类型，并将其设置为 nil。
	Remark      string     `gorm:"type:varchar(255);"`
	CreatTime   *time.Time `gorm:"autoCreateTime;default null"`
	DeletedTime *time.Time `gorm:"default null"`
	UpdateTime  *time.Time `gorm:"autoUpdateTime;default null"`
}

//var db = DBInit()
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
	db.AutoMigrate(&Teacher{})

	return db

}

func main() {
	db := DBInit()

	//now := time.Now()
	//t1 := Teacher{
	//	Tno:    001,
	//	Name:   "游",
	//	Pwd:    "123",
	//	Tel:    "136",
	//	Birth:  &now,
	//	Remark: "第一位老师",
	//}
	////新增单条数据
	//res := db.Create(&t1)
	//fmt.Println(t1)                        // 可以返回修改后的t1 {4 1 游 123 136 2023-04-24 14:15:57.6376609 +0800 CST m=+5.949186501 第一位老师}
	//fmt.Println("错误：", res.Error)          //错误： <nil>
	//fmt.Println("影响行数：", res.RowsAffected) //影响行数： 1
	////创建多条记录
	//t2 := Teacher{
	//	Tno:    002,
	//	Name:   "fei",
	//	Pwd:    "1234",
	//	Tel:    "1360",
	//	Birth:  &now,
	//	Remark: "第二位老师",
	//}
	//teachers := []Teacher{t1, t2}
	//db.Create(&teachers)
	//fmt.Println(teachers)


	//查询
	/*
	teacher :=Teacher{} //实例化一个零值得结构体 代表那张表
	db.First(&teacher)
	fmt.Println("一条查询结果：",teacher) // SELECT * FROM `teachers` ORDER BY `teachers`.`id` LIMIT 1
	teacher1 :=Teacher{}
	db.Last(&teacher1)
	fmt.Println("最后一条查询结果：",teacher1) //SELECT * FROM `teachers` ORDER BY `teachers`.`id` DESC LIMIT 1
	// 检查 ErrRecordNotFound 错误
	//errors.Is(res.Error, gorm.ErrRecordNotFound)
	var teacher2 []Teacher //查所有的需要一个结构体类型的切片 把所有数据返回出来
	db.Find(&teacher2) //SELECT * FROM `teachers`
	fmt.Println(teacher2)
	*/

	var teacher3 Teacher //扫描到结构体里
	db.Model(&Teacher{}).First(&teacher3) //这才是常用的吧
	fmt.Println("db毛豆",teacher3)

	res := map[string]interface{}{} //扫描到map里
	db.Model(&Teacher{}).First(&res)
	fmt.Println(res) //这可以把结果扫描到map里 相当于字典 map[birth:2023-04-24 14:19:17.364 +0800 CST creat_time:<nil> deleted_time:<nil> id:1 my_name:136 name:游 pwd:123 remark:第一位老师 tno:1 update_time:<nil>]

}
