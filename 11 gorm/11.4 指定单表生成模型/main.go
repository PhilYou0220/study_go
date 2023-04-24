package main

import (
"github.com/qmhball/db2gorm/gen"
	"time"
)


type RestrictedArea struct{
	ID uint
	AreaID int
	Name string
	TargetType bool `gorm:"default:0"`
	AreaPoints string
	CenterPoints string
	DateType bool `gorm:"default:0"`
	Deleted bool `gorm:"default:0"`
	PassType bool `gorm:"default:0"`
	ConfigSpeed float64 `gorm:"default:0.0"`
	CreatedBy int `gorm:"default:0"`
	CreatedTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	ConditionType int8 `gorm:"default:0"`
	HasWhiteCar int8 `gorm:"default:0"`
	StopLimit int `gorm:"default:0"`
}


func main(){
	dsn := "epuser:epuser@123-Test@tcp(106.75.138.97:3306)/ep?charset=utf8mb4&parseTime=True&loc=Local"

	//生成指定单表
	tblName := "restricted_area"
	gen.GenerateOne(gen.GenConf{
		Dsn:       dsn,
		WritePath: "./restricted_area",
		Stdout:    false,
		Overwrite: true,
	}, tblName)
}
