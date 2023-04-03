package main

import "fmt"


type Student struct {
	sid int
	name string
	age int
	course []string
}
//直接传地址
func initSid(p *Student){ //结构体也能作为参数传递，p为结构体指针变量
	//(*p).sid=0 //*p 取到结构体的值
	p.sid = 0 //go的语法糖支持把结构体*p写成p
}

func main() {
	//赋值并取地址 &Student可以直接修改原始数据  Student是不行的
	s1 := &Student{sid: 1002, name: "rain", course: []string{"chinese", "math", "english"}}
	fmt.Println(s1.sid) //1002 go的语法糖支持把结构体*p写成p 因此可以看成*s1.sid
	s2 := s1
	s2.sid=1
	fmt.Println(s1.sid) //1
	initSid(s1) //传地址就能改了
	fmt.Println(s1.sid) //0
}
