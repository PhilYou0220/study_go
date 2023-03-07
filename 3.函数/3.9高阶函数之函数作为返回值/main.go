package main

import (
	"fmt"
	"reflect"
)

//Go语言不支持在函数内部声明普通函数，只能声明匿名函数。返回值作为参数那当然是匿名函数了
//函数类型的时候 要强调参数什么类型 返回值是什么类型的函数
func foo() func() int{ //func() int 返回值是一个没有参数 返回int类型的函数
	var inner func() int
	inner = func() int {
		fmt.Println("一个新的函数")
		return 100
	}
	fmt.Println(reflect.TypeOf(inner)) //func() int
	return inner
}

func main() {
	//第一种实现方式 自动判定照片
	a :=foo()()
	fmt.Println(a) //100

	//第二种实现方式
	var b  func() int
	b =foo()
	b()
}
