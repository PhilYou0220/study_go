package main

import "fmt"

func main() {

	//case1
	//foo := func() {
	//	fmt.Println("I am function foo1")
	//}
	//defer foo() //I am function foo1 这里进行了值拷贝 不会因为后面修改了foo而改变
	//foo = func() {
	//	fmt.Println("I am function foo2")
	//}

	//case2
	//x := 10
	//defer func(a int) { //这一拷贝拷贝的是那一刻函数的值和参数的值。注册之后再修改函数值或参数值时，不会生效。
	//	fmt.Println(a)
	//}(x)  // 10
	//x++

	//case3 闭包函数
	x := 10
	defer func() {
		fmt.Println(x)   // 保留x的地址没传参数进来 就去找闭包的x的值
	}() //11
	x++
}
