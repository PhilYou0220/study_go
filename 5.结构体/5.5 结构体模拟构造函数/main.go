package main

import "fmt"

type Student struct {
	sid    int
	name   string
	age    int
	course []string
}

//直接传地址

func NewStudent(sid int, name string, age int, course []string) Student {
	return Student{
		sid:    sid,
		name:   name,
		age:    age,
		course: course,
	}

}

func main() {
	//Go中函数的参数只有位置参数一种,没有关键字参数。
	//好像没啥意思
	s1 := NewStudent(1001,"fei",18,[]string{"1","2"})
	fmt.Println(s1)
}
