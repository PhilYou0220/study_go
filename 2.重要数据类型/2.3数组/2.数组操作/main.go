package main

import (
	"fmt"
	"reflect"
)

func main() {
	//1.索引操作
	var arr = [5]int{1,2,3,4,5}
	//获取值
	fmt.Println(arr[3]) //4
	//修改值
	arr[3]=30
	fmt.Println(arr) //[1 2 3 30 5]
	//2.切片操作 字符串切片完了之后仍是一个字符串 数组切片之后成了一个切片数据
	sl1 := arr[1:3]
	fmt.Println(sl1,reflect.TypeOf(sl1)) //[2 3] []int
	//3.遍历操作，用于容器
	s1 := "iamfei"
	for _,v :=range s1{
		fmt.Println(string(v)) //字符串也可以遍历
	}

	for k,v:=range arr{
		fmt.Println(k,v)
		/*
		0 1
		1 2
		2 3
		3 30
		4 5
		 */
	}

}
