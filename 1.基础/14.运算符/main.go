package main

import "fmt"

func main() {
	const a = 5
	const b = 3
	fmt.Println(5&3) //按位与 相同为1 101 & 011
	fmt.Println(5|3) //按位或 有1为1 101 | 011
	fmt.Println(5^3) //异或 不同为1
	fmt.Println(5<<2) //101整体左移两位后面用0补齐 10100
	fmt.Println(5>>1) //101 整体右移1位后面移出位丢弃 10
}
