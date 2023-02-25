package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//十进制(整型)转二进制(str)
	//v1 := 2
	//r1 := strconv.FormatInt(int64(v1),2) //入参(i int64, base int) 返回 string
	//fmt.Println(r1,reflect.TypeOf(r1)) //10 string

	//十进制转3进制
	//v1 := 80
	//r1 := strconv.FormatInt(int64(v1),3) //入参(i int64, base int) 返回 string
	//fmt.Println(r1,reflect.TypeOf(r1))

	//二进制(str)转十进制(整型)int64
	//bin，要转换的文本
	//2	把文档当做二进制去转换成十进制(整型)
	//16，转换的过程中对结果进行约束
	//结果:如果转换成功，则将err设置为nil,result则永远以int64的类型返回。

	//bin := "1011"
	//r2,err := strconv.ParseInt(bin,2,16)  //(s string, base int, bitSize int) (i int64, err error)
	//fmt.Println(r2,err)

	//练习 十进制14 转 十六进制
	//var a int64 = 14
	//r1 := strconv.FormatInt(a,16)
	//fmt.Println(r1)

	//10011转十进制
	//var text string = "10011"
	//r2,err := strconv.ParseInt( text,2,64)
	//fmt.Println(r2,err)

	////10011转十六进制
	//var text string = "10011"
	//r2,err := strconv.ParseInt( text,2,64)
	//
	//fmt.Println(r2,err)
	//r3 := strconv.FormatInt(r2,16)
	//fmt.Println(r3)
	//var a = 10
	//var b = "1"
	//fmt.Println(a,reflect.TypeOf(a))
	//fmt.Println(b,reflect.TypeOf(b))

	//字符串转浮点
	s := "3.14"
	b, err := strconv.ParseFloat(s, 64)
	fmt.Println(b, err, reflect.TypeOf(b)) //3.14 <nil> float64

	s1 := "12345"
	n1, err := strconv.ParseInt(s1, 10,64)
	fmt.Println(n1, err, reflect.TypeOf(n1)) //12345 <nil> int64

}
