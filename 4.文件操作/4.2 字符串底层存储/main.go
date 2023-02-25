package main

import "fmt"

func main() {
	//字符串是双引号 go语⾔的string是⼀种数据类型，这个数据类型占⽤16字节空间，前8字节是⼀个指针，指向字符串值的地址，后⼋个字节是⼀个整数，标识字符串的长度；和切片有点像
	//字符串底层byte类型的数组
	//字符串不能被修改
	var a string
	a = "123hello"
	//a[0]=3  cannot assign to a[0]
	fmt.Println(a[0]) //49
	fmt.Println(a[1:4]) //23h
}
