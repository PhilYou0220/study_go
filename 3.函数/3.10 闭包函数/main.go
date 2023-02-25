package main

import (
	"fmt"
	"time"
)

//和python差不多三个条件 1.函数里有函数（go当然是匿名函数） 2.内函数使用了外部变量 3.外部函数的返回值是内函数
//编码遵循封闭开放原则 封闭原则：函数一旦写好了就不要去动了

func getTimer(f func()) func() {
	counter := func()  { //不需要传参初始化的时候就定义了参数
		start := time.Now().Unix()
		f()
		end := time.Now().Unix()
		fmt.Printf("消耗时间为：%d",end-start)

	}
	return counter
}




func bar(){
	fmt.Println("我执行了")
	time.Sleep(time.Second*3)
}
func main() {
	bar := getTimer(bar) //开放原则 不影响其他功能的使用所以命名为bar
	bar()
}
