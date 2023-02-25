package main

import "fmt"

func main() {
	//打印个空切片
	var s []int
	//fmt.Println(s) //[] nil吗
	//
	s1 := append(s,1)
	fmt.Println(s1,len(s1),cap(s1))
	s2 := append(s,2,3,4 )
	fmt.Println(s1,len(s1),cap(s1))
	fmt.Println(s2,len(s2),cap(s2))
	////追加另外一个切片
	//var t =[]int{5,6,7}
	//s3:= append(s1,t...)
	//fmt.Println(s3)

	//都是基于底层数组操作的
	//var s4 = make([]int,3,10)
	//fmt.Println(s4) //[0 0 0]
	//s5 := append(s4,100)
	//fmt.Println(s5) //[0 0 0 100]
}
