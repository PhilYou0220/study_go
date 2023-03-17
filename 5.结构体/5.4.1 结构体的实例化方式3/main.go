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
	p.sid = 0 //go的语法糖支持把*p写成p
}

func main() {
	s1 := Student{sid: 1002, name: "rain", course: []string{"chinese", "math", "english"}}
	fmt.Println(s1.sid) //1002
	s2 := s1
	s2.sid=1
	s2.course[1]="a"
	fmt.Println(s1.course) //都会被修改
	fmt.Println(s2.course) //会被修改 切片是引用类型

	initSid(&s1) //传地址就能改了
	fmt.Println(s1.sid) //0
	fmt.Println(s2.sid)

}
