package main

import "fmt"

func main() {
	//fmt.Printf可以打印出拼接的字符串
	e, f, g := "a", "b", "c"
	fmt.Printf("第一个%s,第二个%s，第三个%s", e, f, g)
	fmt.Println("请输入：")
	var (
		a string
		b string
		c string
	)
	_, _ = fmt.Scanln(&a, &b, &c)
	// fmt.Sprintf可以使用变量接收
	result := fmt.Sprintf("第一个%s,第二个%s，第三个%s", a, b, c)
	fmt.Println(result)
	result1 := "第一个" + a + "第二个" + b + "第三个" + c
	fmt.Println(result1)
}
