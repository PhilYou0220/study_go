package main

import "fmt"

func main() {
	//var x int
	//var y int
	//
	//x = 10
	//// y开辟一个新的地址空间 并拷贝x的值
	//y = x
	//fmt.Println(x,y)

	//c语言经典例题，交换两个值
	//var x int = 10
	//var y int = 20
	////go语言提供的方法
	////x,y = y,x
	//
	////经典方法
	//t := 10
	//x = y
	//y = t
	//
	//fmt.Println(x,y)
	var a,_ = 10,2
	fmt.Println(a,_) // cannot use _ as value

}
