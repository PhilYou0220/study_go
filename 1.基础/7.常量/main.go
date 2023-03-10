package main

import "fmt"

// 常量一般放在顶部作为全局常量




func main() {
	//常量不可被修改
	//const name int = 3
	//const age = 18
	//age = 19  //  cannot assign to age


	//常量必须赋值
	//const name string //const declaration cannot have type without expression
	//name = "3"

	// 支持批量定义
	const (
		va = 1
		la string  = "1"
		vi bool = false
	)
	fmt.Println(va,la,vi)

	// iota(默认为0)
	const (
		v1 = iota+100
		_
		v2
	)
	fmt.Println(v1,v2)  // 100 102
}
