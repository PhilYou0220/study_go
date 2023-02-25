package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//不同整型之间不能相加减
	//a := 16 // 这样定义直接是int型
	//fmt.Println(reflect.TypeOf(a)) //int
	//var b int16 = 16
	////fmt.Println(a+b) //invalid operation: a + b (mismatched types int and int16)

	//整型转换
	//var a int8 = 10
	//var b int16 = 20
	//fmt.Println(int16(a)+b)

	//var a int8 = 10
	//var b int16 = 129
	//fmt.Println(int8(b)) //-127

	//整型转字符串
	//var a int = 16
	//b := strconv.Itoa(a)
	//fmt.Println(b,reflect.TypeOf(b)) // 16 string
	//const c string = "1"
	//fmt.Println(reflect.TypeOf(c)) //string

	//字符串转整型
	var name string = "phil"
	b,err := strconv.Atoi(name)
	if err ==nil{
		fmt.Println(b,reflect.TypeOf(b)) // 666 int
	}else {
		fmt.Println(err) // strconv.Atoi: parsing "phil": invalid syntax
	}

}
