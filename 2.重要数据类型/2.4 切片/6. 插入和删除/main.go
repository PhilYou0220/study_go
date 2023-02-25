package main

import "fmt"

func main() {
	//切片的 使用同一个变量 长度和容量重新写入
	var s = []int{1, 2, 3}
	fmt.Printf("%p \n", &s)        //0xc000096440
	fmt.Println(s, len(s), cap(s)) //[1 2 3] 3 3
	s = append(s, 4)
	fmt.Printf("%p\n", &s)         //0xc000096440
	fmt.Println(s, len(s), cap(s)) //[1 2 3 4] 4 6

	//头部插入值或切片
	var s1 = []int{1, 2, 3}
	s1 = append([]int{0}, s1...) //[]int{0}只是没有变量指向它而已
	fmt.Println(s1)              //[0 1 2 3]

	//任意位置插入元素 切成两段进行添加
	var b = []int{1, 2, 3, 4, 5}
	var i = 2
	fmt.Println(append(b[:i], append([]int{100}, b[i:]...)...)) //[1 2 100 3 4 5]

	//删除索引 2
	fmt.Println(append(b[:i],b[i:]...)) //[1 2 3 4 5]

}
