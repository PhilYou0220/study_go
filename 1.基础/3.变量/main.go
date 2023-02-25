package main

import "fmt"

func main() {
	//声明+赋值
	//var sd string = "phil"
	//fmt.Println(sd)
	//
	//var num int = 7
	//fmt.Println(num)
	//
	//var flag = true
	//fmt.Println(flag)
	////先声明再赋值
	//var job string
	//job = "null"
	//fmt.Println(job)

	var (
		name = "武沛齐"
		age= 18
		hobby ="大保健"
		salary = 1000000
		gender string//只声明但不赋值，有一个默认:"
		length int //只声明但不赋值，有一个默认:0
		sb bool //只声明但不赋值，有一个默认:false
		)
	fmt.Println(name,age,hobby,salary,gender , length, sb)
}
