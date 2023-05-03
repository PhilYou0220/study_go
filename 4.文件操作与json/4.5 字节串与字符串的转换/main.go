package main

import "fmt"

func main() {
	//字符串转字节串[]byte(a) 优点像强制类型转换
	var  a string
	a  = "飞abc"
	b := []byte(a)
	fmt.Println(b) //[233 163 158 97 98 99]

	//字节串转字符串string
	c := []byte{233,163,158,97,98,99}
	fmt.Println(string(c)) //飞abc
}
