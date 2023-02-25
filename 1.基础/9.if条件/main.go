package main

import "fmt"

func main() {
	//go语言中 and 逻辑与 使用 &&
	const n1 = 1
	const  n2 int = 2
	const  n3 = 3
	if n2>n1 && n3>=n1{
		fmt.Println("很好")
	}
	//go语言中 or 逻辑或 使用 ||
	if n2>n3 || n2 >n1{
		fmt.Println("good")
	}
	//go语言中 not 逻辑非 使用 ！
	flag := false
	//if ! flag{
	//	fmt.Println("perfect")
	//}



}
