//uber 开源的日志库 强类型日志
package main

import "go.uber.org/zap"

func main(){
	//获取logger对象
	//zap.NewExample()
	//zap.NewDevelopment()
	log, _ := zap.NewProduction()

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
	log.Info("日志信息",zap.String("name",name),
		zap.Ints("list",list),
		zap.Int64("age",age),
		zap.Any("any",nameList), //任意类型
	//{"level":"info","ts":1684843048.9437547,"caller":"12 zap日志库/main.go:27","msg":"日志信息","name":"张三","list":[0,1,2,3,4,5,6,7,8,9],"age":16,"any":["李四","","","",""]}
		)
}
