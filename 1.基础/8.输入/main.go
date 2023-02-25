package main

import "fmt"

func main() {
	//fmt.Scan输入一个值
	//fmt.Print("请输入姓名：")
	//var name string
	//// n 输入的个数，err错误信息，如果没有错误 返回空指针 nil
	//n,err :=fmt.Scan(&name)  // 接收输入的信息 以空格分割
	//fmt.Println(name)
	//fmt.Println(n,err)  // 1 <nil>

	////fmt.Scan输入多个值,输入两（多）个值时必须输入多个值 否则会一直等待
	//fmt.Print("请输入姓名：")
	//var (
	//	name string
	//	age int
	//)
	//
	//n,err :=fmt.Scan(&name,&age)
	//fmt.Println(name,age)
	//fmt.Println(n,err)

	//fmt.Scanln() 用法基本一致 是在等待回车
	//fmt.Print("请输入姓名：")
	//var (
	//	name string
	//	age int
	//)
	//
	//n,err :=fmt.Scanln(&name,&age)
	//fmt.Println(name,age)
	//fmt.Println(n,err)

	//fmt.Scanf() 用法基本一致 是在等待回车 但是需要输入模板才行 如 我叫%s 今年%d 岁 提取关键位置参数
	fmt.Print("请输入姓名：")
	var (
		name string
		age int
	)
	
	n,err :=fmt.Scanf("我叫%s 今年%d 岁",&name,&age)
	fmt.Println(name,age)
	fmt.Println(n,err)

}
