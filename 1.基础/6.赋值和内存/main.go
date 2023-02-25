package main

import "fmt"
//本节示例int,string,和bool 都是值引用
func main() {
	//
	name := "武沛齐"
	nickname := name
	fmt.Println(name,&name) // 武沛齐 0xc0000881e0
	fmt.Println(nickname,&nickname) // 武沛齐 0xc0000881f0

	name = "alex"
	fmt.Println(name,&name) // alex 0xc0000881e0
	fmt.Println(nickname,&nickname) // 武沛齐 0xc0000881f0

}
