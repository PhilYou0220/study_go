package main

import (
	"fmt"
	"time"
)
//高阶函数 之 参数是个函数
func timer(f func()) {
	start := time.Now().Unix()
	f()
	end := time.Now().Unix()
	fmt.Println("执行花费时间",end-start)
}

func bar(){
	fmt.Println("bar函数开始执行")
	time.Sleep(time.Second*2)
	fmt.Println("bar函数结束执行")
}

func main() {
	timer(bar)

}
