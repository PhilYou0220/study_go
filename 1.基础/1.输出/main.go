package main

import "fmt"

func main() {
	//print("内置print")
	//println("内置println")
	fmt.Print("print")
	fmt.Println("println")
	fmt.Printf("老汉开着%s，去接alex这个%d货。\n", "兰博基尼", 2)

	//格式化字符串
	a := 1
	b := "张三"
	fmt.Printf("他是%s,%d岁了", b, a) //他是张三,1岁了

	//有个返回值 可以进行其它操作
	s := fmt.Sprintf("他是%s,%d岁了", b, a)
	fmt.Println(s)

	//多行字符串
	s1 := `第一行
			第二行
			第三行
			`
	fmt.Println(s1)

}
