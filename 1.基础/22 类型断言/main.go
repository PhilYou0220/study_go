//由于接口类型的值什么类型都有，因此可以断言是否是需要的值
package main

import "fmt"

func main(){
	// 定义一个空接口
	var x interface {}
	x = 100
	//猜测x是否是bool值，如果猜对了就把接口变量转为对应的类型
	//猜错了b 默认猜测类型的零值
	//b,ok :=x.(bool)
	//fmt.Println(b,ok) //false false

	b,ok :=x.(int)
	fmt.Println(b,ok) //100 true

	//type y struct {
	//	uid int64
	//}
	//
	//c :=y.(int64)
	//fmt.Println(c)
}