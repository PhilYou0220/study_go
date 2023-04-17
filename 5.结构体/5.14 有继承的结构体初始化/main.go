package main

import "fmt"
//会直接初始化几次继承后的结构体
type Addr struct {
	province string
	city     string
	county   string
}

type Person struct {
	name string
	age  int
	Addr //匿名字段 结构体嵌套 相当于 Addr Addr
}

func main()  {
	fmt.Printf("%+v",&Person{})
	fmt.Printf("%p",&Person{})
}