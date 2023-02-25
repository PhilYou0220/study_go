package main

import "fmt"

func main() {
	// 整型
	fmt.Println(666)
	fmt.Println(6 + 9)
	fmt.Println(6 - 9)
	fmt.Println(6 * 9)
	fmt.Println(16 / 9) // 取商 相当于取证 由于golang是静态语言这会得到整数
	fmt.Println(16 % 9) //fmt.Println(16//9) 没有python中的取整

	//字符串类型，特点：通过双引号 不像python可以单引号
	fmt.Println("phil游")
	fmt.Println("phil" + "666") // 字符串拼接 和python一样直接拼接
	//fmt.Println("phil"+ 666) 混合数据类型无法拼接

	// bool类型 得到 true or false
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println("a"=="A")
}
