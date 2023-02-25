package main

import "fmt"

type Person struct {
	name string
	age  int
	Addr //匿名字段 结构体嵌套 相当于 Addr Addr
}

type Addr struct {
	province string
	city     string
	county   string
}

func (a *Addr) history() {
	fmt.Println("这是一座历史悠久的城市")
}
func main() {
	p1 := Person{"fei", 18, Addr{"四川", "遂宁", "射洪"}}
	fmt.Println(p1)
	p1.Addr.history() //这是一座历史悠久的城市
	p1.history() // 可以直接调用

}
