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
	//new(结构体)
	//s1 := new(Student{sid: 1002, name: "rain", course: []string{"chinese", "math", "english"}}) //这样不行
	s1 := new(Student) //相当于 &Student{}
	fmt.Println(s1.sid)
	s2 := s1
	s2.sid=1
	initSid(s1)
	fmt.Println(s1.sid)

	//new 与make的区别 new 返回的是个指针（地址） make 返回的是个 就数据类型
	//但是make只用于slice、map以及channel的初始化（非零值）；而new用于类型的内存分配，并且内存置为零
}
