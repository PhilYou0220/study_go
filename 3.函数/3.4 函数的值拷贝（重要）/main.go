package main

import "fmt"

//func foo(x int){
//	x=100
//}
//传地址发生值拷贝的时候拷贝的是地址

func foo1(x *int){
	*x=100
}
func main() {
	//var a = 1
	////foo(1) //如果直接传1 相当于 函数 x =1 发生的是赋值操作
	//foo(a) // 如果是传x 就发生了 值拷贝 foo.a 与 main.x 是相互隔离的
	//fmt.Println(x) //1

	var p *int
	a := 1
	p=&a
	foo1(p)
	fmt.Println(a)


}
