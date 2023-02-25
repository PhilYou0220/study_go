package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 字符都是单引号 byte
	//byte 只能是ascii的那256的数字
	var x byte //go语言中byte就是unit8(0-255) 8bit 所以是值类型
	x = 'a'
	fmt.Println(x, reflect.TypeOf(x)) //97 uint8

	fmt.Printf("%c\n", x) //a c:char 显示字符
	fmt.Printf("%d\n", x) //97 10进制
	fmt.Printf("%b\n", x) //1100001 2进制
	fmt.Printf("%x\n", x) //61 16进制

	var b2 uint8 // 等价于 byte
	b2 = 65
	// b2 = 'c'
	fmt.Printf("%c\n",b2) //A
	fmt.Printf("%d\n",b2) //65
	fmt.Println(b2) // ASCII数字65
}
