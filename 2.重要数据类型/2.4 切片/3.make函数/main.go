package main

import "fmt"

func main() {
	//引用类型都需要开辟空间,否则赋值会报错
	//var sl1 []int
	//fmt.Println(sl1,reflect.TypeOf(sl1))
	//sl1[0]=1 //Indexing may panic because of 'nil' slice

	//make函数 make([]Type, size, cap) 声明并开辟空间
	sl1 :=make([]int,5) //切片的底层是数组 int 类型的默认值是 0 cap省略的话就默认为长度
	fmt.Println(sl1,len(sl1),cap(sl1)) //[0 0 0 0 0] 5 5
	var s = make([]int,5,10) //分配内存，扩展空间没有
	fmt.Println(s)
	s[4]=100
	fmt.Println(s) //[0 0 0 0 100]

}
