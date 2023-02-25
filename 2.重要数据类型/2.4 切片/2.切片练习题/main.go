package main

import "fmt"

func main() {
	//第一题
	//s1 := []int{1,2,3}
	//s2 :=s1[1:] //[2 3]
	//s2[1]=4 //[2 4]
	//fmt.Println(s1) //[1 2 4]

	//第二题
	var a =[]int{1,2,3} //底层 数组 [1 2 3]
	b:=a
	a[0]=100
	fmt.Println(b)
}
