package main

import "fmt"
//结构体是值类型 不是引用类型
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
	fmt.Println(Student{}) //{0  0 []} 存在默认值
}
