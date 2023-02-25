package main

import (
	"fmt"
	"reflect"
)

func test(x ...int){ // 这是一个int类型的切片
	fmt.Println(x,reflect.TypeOf(x))

}

func main() {

	test(1,2,3) //[1 2 3] []int
}
