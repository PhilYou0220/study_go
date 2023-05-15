package main

import (
	"fmt"
	"reflect"
	"time"
)

//type Todo struct {
//Title  string
//Status bool
//}

func main() {
	//var todo Todo
	//todo =Todo{}
	//fmt.Println(todo)
	//auth_param := "bear 23123 132"
	//part := strings.Split(auth_param, " ")
	//fmt.Println(part)
	//for k,v:= range part{
	//	fmt.Println(k,v)
	//}
	fmt.Println(time.Now()) //2023-05-13 12:30:44.5024001 +0800 CST m=+0.004532601
	fmt.Println(reflect.TypeOf(time.Now()))
	fmt.Println(time.Now().Add(time.Second*10).Unix()) //1683952254 转时间戳
}
