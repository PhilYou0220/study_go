package main

import "fmt"

type Student struct {
	name string
	age int
}

func main() {
	fmt.Println(&Student{}) //这是结构体的地址 只是打印出来是&{ 0},如果要看地址使用%
	name := "za"
	fmt.Println(&name)//0xc00003e260

	fmt.Printf("%p",&Student{}) //0xc0000040a8
}
