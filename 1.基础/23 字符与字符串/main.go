package main

import "fmt"

func main(){
	//1.uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
	//2.rune类型，代表一个UTF-8字符。
	//要用到 rune 类型。当需要处理中文、日文或者其他复合字符时
	//rune 类型实际是一个 int32 8bit一个字节 int32 4个字节


	s0 := '中' //这是一个字符类型 rune  相当于utf-8编码 不会被转义 记住有什么转义问题就用 rune类型
	s1 := "中" //这是一个字符串类型
	s2 := "abc游飞"
	fmt.Println(s0) //20013
	fmt.Println(s1) //中
	fmt.Println([]rune(s2)) //
	fmt.Println([]byte(s2))
	for k,v := range s2{  //rune 直接记住for range 就是遍历字符
		fmt.Println(k,string(v))
	}
}
