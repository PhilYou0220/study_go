package main

import (
	"fmt"
	"golang.org/x/crypto/ssh/agent"
)

//func main() {
//	name :="游"
//	fmt.Println(name)
//
//	if true {
//		age := 18
//		name := "phil"
//		fmt.Println(age)
//		fmt.Println(name)
//	}
//	fmt.Println(name)
//}

//全局变量(不能以省略的方式)
var school string ="老男孩IT教育"//可以
//var school = //可以
//school := "老男孩T教育" //不可以
var (
	v1 = 123
	v2 ="你好"
	v3 int
)
func main() {
	name := "武沛齐" //局部变量fmt .println(name)
	if true {
		age := 18//局部变量
		// name := "alex"//局部变量
		fmt.Println(age)
		fmt.Println(name)
	}
	//fmt.Println(age)
	age := 20 //在同一个函数中在不同的作用域内定义相同名称的变量是可以的 比如age在 if条件里的age和外部的age互不相关
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(school)
	fmt.Println(v1, v2, v3)
}