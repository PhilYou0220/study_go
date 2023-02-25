package main

import "fmt"

type Student struct {
	sid    int
	name   string
	age    int
	course []string
}

func initSid(s Student) { //结构体也能作为参数传递
	s.sid = 0
}

func main() {

	s1 := Student{sid: 1002, name: "rain", course: []string{"chinese", "math", "english"}}
	fmt.Println(s1.sid) //1002
	s2 := s1
	s2.sid = 1
	initSid(s1)         //此处发生了值拷贝 s与s2是两块空间
	fmt.Println(s1.sid) //1002

}
