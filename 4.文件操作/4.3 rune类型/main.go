package main

import "fmt"

func main() {

	// type rune = int32 所以是值类型,32位4字节,但是它表示的是一个unicode符号,但是注意unicode编码是两个字节
	var a rune
	//a = '张三' //invalid character literal (more than one character)
	a = '飞'
	fmt.Println(a,string(a)) //39134 飞

}
