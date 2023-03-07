package main

import "fmt"
//记住函数传参是拷贝值 切片扩容后切片的地址变了 切片存的地址也变了
func  foo(a []int)  {
	a[0]=100
	a = append(a, 4)
	a[1]=200
	fmt.Println(a,len(a),cap(a))//[100 200 3 4] 4 6
}

func main() {
	var a = []int{1,2,3}
	foo(a) //整型切片引用类型
	fmt.Println(a) //[100 2 3]
}
