package main

import "fmt"

type Person struct {
	name string
	age  int
	Addr //匿名字段 结构体嵌套 相当于 Addr Addr
}

type Addr struct {
	province string
	city string
	county string
}

func main() {
	p1 := Person{"fei", 18,Addr{"四川","遂宁","射洪"}}
	fmt.Print(p1)
	fmt.Println(p1.Addr.county) //可以通过匿名字段去取
	fmt.Println(p1.county) //也可以直接取
}
