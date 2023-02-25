package main

import "fmt"

func main() {
	//基本数据类型都属于是值类型
	//值类型（整型、浮点型、字符串、布尔型、数组、结构体）：声明未赋值之前存在默认值
	var a int
	fmt.Println(&a)

	//引用类型包括指针类型、切片、map、channel
	//引用类型声明未赋值之前没有开辟空间，即没有默认值(nil)
	//var p *int
	//*p =10 //anic: runtime error: invalid memory address or nil pointer dereference
	var p *int
	p = new(int) //给p存的地址开辟一块空间,默认值为0的地址
	fmt.Println(*p)
	//fmt.Println(&p)
	*p=10
	fmt.Println(*p)


}
