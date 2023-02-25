package main

import "fmt"

func main() {
	//声明+赋值
	var sd string = "phil"
	fmt.Println(sd)

	var num int = 7
	fmt.Println(num)

	var flag = true
	fmt.Println(flag)
	//先声明再赋值
	var job string
	job = "null"
	fmt.Println(job)

	//变量简写
	name := "wo"
	fmt.Println(name)
}
