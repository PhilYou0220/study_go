package main

import "fmt"

//类型包含属性和方法
//对象是类型的实例化结果
//声明结构体  type 类型名 struct，属性名不能重复
type Student struct {
	sid int
	name string
	age int
	course []string
}
func main() {
	var s Student //声明一个结构体对象 ,值类型，默认开辟空间，字段赋予零值 或者叫实例化
	fmt.Println(s)
	//赋值
	s.sid =1001
	s.name = "张三"
	s.age  =18
	s.course = make([]string,3) //切片需要开辟空间
	s.course[0] = "chinese"
	fmt.Println(s)

	//结构体 是值类型，因此它的地址是第一个值的地址，后面的变量地址是连续的
	fmt.Printf("%p\n",&s) //0xc000058040
	fmt.Println(&s.sid) //0xc000058040
	fmt.Println(&s.name) //0xc000058048
	fmt.Println(&s.age) //0xc000058058
	fmt.Printf("%p",&(s.course)) //0xc000058060
}
