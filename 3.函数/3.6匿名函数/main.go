package main

import (
	"fmt"
	"reflect"
)

func foo(){
	//Go语言不支持在函数内部声明普通函数，只能声明匿名函数。
	var bar func(int,int) int//声明一个参数为int,int返回值为int的函数类型
	bar = func (x int,y int)int {
		return 100
	}
	fmt.Println(reflect.TypeOf(bar))
}

func main() {
	//匿名函数，顾名思义，没有函数名的函数
	//1.直接调用匿名函数
	//(func() {
	//	fmt.Println("hello world")
	//})()


	(func(x, y int) {
		fmt.Println(x + y)
	})(1,2)

	foo() //func(int, int) int
}
