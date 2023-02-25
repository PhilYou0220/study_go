package main

import (
	"fmt"
	"reflect"
)

func main() {
	//特性 1.一致性，数组内元素类型一致 2.有序性 数组中的元素是有序的，存储在一块连续内存，通过下标访问。3 长度不可变性：数组一旦初始化，则长度（数组中元素的个数）不可变
	//var 数组名 [元素数量]元素类型

	//0.数组是值类型，存在默认值 存储在一块连续内存
	var v1 [3]int
	fmt.Println(v1)	//[0 0 0]
	fmt.Println(&v1[0],&v1[1],&v1[2]) //0xc0000123a0 0xc0000123a8 0xc0000123b0
	var v2 [3]*int
	fmt.Println(v2) //[<nil> <nil> <nil>]

	//1.数组声明 再赋值
	var a [3]int
	a[0],a[1],a[2] = 1,2,3
	fmt.Println(a,reflect.TypeOf(a)) // [1 2 3] [3]int
	//2.声明并赋值
	var b = [3]string{"i","am" ,"fei"}
	fmt.Println(b) //[i am fei]
	//3.不限长度的数组 go编译器进行判断
	var c = [...]int{1,23,45,4 }
	fmt.Println(c,len(c)) //[1 23 45 4] 4
	//4.数组可以通过索引赋值
	var d = [3]int{0:1,2:111}
	fmt.Println(d) //[1 0 111]
}
