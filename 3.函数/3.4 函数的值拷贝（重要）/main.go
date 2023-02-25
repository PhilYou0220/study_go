package main

import "fmt"

//func foo(x int){
//	x=100
//}


func foo1(x *int){
	*x=100
}
func main() {
	//var x = 1
	////foo(1) //如果直接传1 相当于 函数 x =1 发生的是赋值操作
	//foo(x) // 如果是传x 就发生了 值拷贝 foo.x 与 main.x 是相互隔离的
	//fmt.Println(x) //1

	var p *int
	a := 1
	p=&a
	foo1(p)
	fmt.Println(a)


}
