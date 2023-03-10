package main

import (
	"fmt"
)

func main() {
	//切片是对数组的重复引用 引用类型特点：变量通过存储一个地址来存储最终的值。内存通常在堆上分配，通过GC回收。
	//切片构建方式一：通过数组构建切片
	//var arr = [5]int{1, 2, 3, 4, 5}
	//s1 := arr[1:]   //  s1存了三个东西 起始地址 长度(切片索引end-start) 容量(原数组长度-原数组start索引)
	//fmt.Println(s1) //[2,3,4,5]
	//s2 := s1[1:3]   //基于s1切片再次切片  长度 和容量 都是基于数组进行计算的
	//fmt.Println(s2) //s2 = [3,5]
	//fmt.Println(s1,reflect.TypeOf(s1)) //[2 3 4 5] []int

	//方式2 直接声明切片 与数组声明比较像
	var sl1 = []int{10, 11, 12, 13, 14}
	//底层相当于创建了这个数组 然后再切片
	//arr :=[5]int{10, 11, 12, 13, 14}
	//sl1 := arr[:]
	s1 := sl1[1:4] //[11, 12, 13]
	fmt.Println(len(s1),cap(s1)) //3 4
	s2 := s1[2:] //[13]
	fmt.Println(len(s2),cap(s2)) // 1 3 错了 是1 2 容量 是原数组长度-原数组start索引  13在原数组的索引是3 5-3=2
}
