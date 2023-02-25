package main

import (
	"fmt"
	"reflect"
)

func foo(){
	fmt.Println("foo")
}


func main() {
	fmt.Println(foo,reflect.TypeOf(foo)) //0x49ddf0 func()
	
}
