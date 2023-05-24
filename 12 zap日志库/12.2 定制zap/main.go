package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
	"os"
)

//可以参考https://www.liwenzhou.com/posts/Go/zap/


func main()  {
	//zapcore.Core需要三个配置——Encoder(编码器)，WriteSyncer（写到哪里），LogLevel（日志级别）。
	//1、encoder
	encoder :=zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()) //zap.NewProductionEncoderConfig()一些json
	//2、WriteSyncer（写到哪里）
	//新建文件
	file, _ := os.OpenFile("./test.log",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
	//传入文件句柄输出到文件里
	writeSyncer :=zapcore.AddSync(file)
	//输出到终端里
	//consoleSyncer:=zapcore.AddSync(os.Stdout)

	//3、设置日志级别
	//创建zapcore
	core :=zapcore.NewCore(encoder,writeSyncer,zapcore.InfoLevel)
	/*
	支持从配置文件读取
	s := "info"
	level,err:=zapcore.ParseLevel(s)
	if err!=nil{
		level =zapcore.InfoLevel
	}
	*/
	//多核心 core :=zapcore.NewCore(encoder,zapcore.NewMultiWriteSyncer(writeSyncer,consoleSyncer),zapcore.InfoLevel)
	logger:=zap.New(core)
	var (
		name string
		age  int64
		list  []int
	)
	for i := 0; i < 10; i++ {
		list = append(list, i)
	}

	nameList :=make([]string,5)
	nameList[0] = "李四"
	age =16
	name = "张三"

	//需要传个字符串，然后zap.type指定参数的类型，这也是他快的原因
	logger.Info("日志信息",zap.String("name",name),
		zap.Ints("list",list),//int类型的切片
		zap.Int64("age",age),
		zap.Any("any",nameList), //任意类型
		//{"level":"info","ts":1684843048.9437547,"caller":"12 zap日志库/main.go:27","msg":"日志信息","name":"张三","list":[0,1,2,3,4,5,6,7,8,9],"age":16,"any":["李四","","","",""]}
	)

}