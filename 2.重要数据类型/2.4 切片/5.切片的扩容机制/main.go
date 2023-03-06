package main

import "fmt"

func main() {

	//扩容
	a := []int{11,22,33}
	fmt.Println(len(a),cap(a))//3,3
	//c:=append(a,44) //出现扩容，双倍扩容 生成一个新的数组
	fmt.Printf("%p \n",a)   // 直接打印是第一个切片的内存地址
	fmt.Println(&a[0]) // 第一个元素的的的内存地址
	fmt.Printf("%p \n",&a[0]) // 第一个元素的的的内存地址
	fmt.Printf("%p \n",&a) //切片的地址


	//a[0]=100
	//fmt.Println(a) //[100 22 33]
	//fmt.Println(c,len(c),cap(c),&c[0]) // [11 22 33 44] 4 6

	// 练习题
	//a := make([]int,3,10)
	//fmt.Println(a) //[0 0 0]
	//b:=append(a,11,22) // 起始地址 长度5 容量10
	//
	//fmt.Println(a) //a的切片 长度存的3那就是3 扩容了也是3 [0 0 0]
	//
	//fmt.Println(b) //[0 0 0 11 22]
	//a[0]=100
	//fmt.Println(a) //[100 0 0]
	//fmt.Println(b)  //[100 0 0 11 22 ]

	arr := [4]int{10, 20, 30, 40}
	s1 := arr[0:2] //[10 20] &10 2 4 长度为2 只用了 10 和 20
	s2:= s1 //拷贝了一份切片三要素
	s3 := append(append(append(s1,1),2),3) //底层数组arr 双倍扩容 出现新的数组
	s1[0]=1000
	fmt.Println(s1) //[1000 20]
	fmt.Println(s2) // [1000 20]
	fmt.Println(s3) // [10 20 1 2 3 ]
	fmt.Println(arr) //[1000 20 1 2]

}
