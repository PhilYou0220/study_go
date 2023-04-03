package main

import (
	"fmt"
	"reflect"
)

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

	// iota(默认为0) 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1，当遇到被赋予其他值时会直接跳过
	const (
		v1 = iota+100
		_
		v2
	)
	fmt.Println(v1,v2)  // 100 102


	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a,b,c,reflect.TypeOf(a)) //0 1 2 int
}
