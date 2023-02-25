package main

import "fmt"

type Student struct {
	sid    int
	name   string
	age    int
	course []string
}
//这是方法不是函数 和python很像 类里的叫方法 类外的叫函数
func (s Student)read(){
	fmt.Printf("我在读书\n")
}

func(s Student)learn(){
	//可以直接使用结构体中的字段
	fmt.Printf("%s在学习",s.name)
}



func main() {
	//Go中函数的参数只有位置参数一种,没有关键字参数。
	s1 := Student{sid: 1001,name: "fei",age: 18,course: []string{"chinese", "math", "english"}}
	// 实例化之后可以直接调用
	s1.read() //发生了值拷贝 s = s1 函数声明了 所以不需要声明

	s1.learn() //发生了值拷贝 s = s1  所以能调用

}
