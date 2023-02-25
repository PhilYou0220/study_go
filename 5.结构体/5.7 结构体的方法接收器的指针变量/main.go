package main

import "fmt"

type Student struct {
	sid    int
	name   string
	age    int
	course []string
}
//这是方法不是函数 和python很像 类里的叫方法 类外的叫函数
func (s *Student)addAge(){
	s.age += 1
	fmt.Printf("年龄增加了，现在%d\n",s.age)
}

func(s *Student)downAge(){
	//可以直接使用结构体中的字段
	s.age -= 1
	fmt.Printf("年龄减少了了，现在%d\n",s.age)
}



func main() {
	//Go中函数的参数只有位置参数一种,没有关键字参数。
	s1 := Student{sid: 1001,name: "fei",age: 18,course: []string{"chinese", "math", "english"}} //当然写成 &Student也不会错
	// 实例化之后可以直接调用
	s1.addAge() //由于参数是指针 会发生 s = &s1

	s1.downAge()

}
