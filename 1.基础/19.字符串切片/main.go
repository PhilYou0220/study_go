package main

import (
	"fmt"
	"reflect"
)

func main() {
	s:="international"
	s1 := "张三"
	// 字符串取索引  与python左闭右开区间相同 但是不支持负索引 都是从0开始计数
	fmt.Println(string(s[1])) //n
	fmt.Println(string(s1[1])) //取中文存在问题 中文是由三个字节组成

	//字符串切片
	fmt.Println(string(s[0:3]),reflect.TypeOf(string(s[0:3]))) //int
	fmt.Println(string(s1[:])) //取所有
	fmt.Println(string(s1[0:3])) //张 取前三个 就是张 utf-8中中文是由三个字节组成

	//字符串拼接
	a := "ph"
	b := "il"

	fmt.Println(a + b)
	fmt.Printf("%s%s \n",a,b)

	//转义符 对于 \ "  等特殊字符进行转义
	fmt.Println("\\") // \
	fmt.Println("\"") // "

	//打印多行字符串 使用` 反引号 python是""" """
	info := `
		1.进攻
		2.防守
		3.撤退
	`
	fmt.Println(info)


}
