//package main
//
//import (
//	"fmt"
//	"reflect"
//	"unsafe"
//)
//
//func main() {
//	//第一题
//	//s1 := []int{1,2,3}
//	//s2 :=s1[1:] //[2 3]
//	//s2[1]=4 //[2 4]
//	//fmt.Println(s1) //[1 2 4]
//
//	//第二题
//	var a =[]int{1,2,3} //底层 数组 [1 2 3]
//	b:=a
//	fmt.Println(reflect.TypeOf(b))
//	//fmt.Println(reflect.TypeOf(b))
//	a[0]=100
//	fmt.Println(b)
//
//
//}

package main

import (
"fmt"


"reflect"
"unsafe"
)

func main() {

	//扩容
	b := [3]int{11, 22, 33}
	a := b[:]
	fmt.Printf("%p %p \n", a, &b) //
	fmt.Println(&a[0])
	fmt.Printf("%v\n", unsafe.Pointer(&a)) //adress of slice
	c := b
	fmt.Println(reflect.TypeOf(c))
	rc := reflect.TypeOf(c)
	ra := reflect.TypeOf(a)
	fmt.Println(rc.Kind(), ra.Kind())
	fmt.Printf("%v\n", unsafe.Pointer(&c)) // adress of slice
	fmt.Printf("%p %p\n", &a, &c)
}