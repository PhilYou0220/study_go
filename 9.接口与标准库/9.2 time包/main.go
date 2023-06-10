package main

import (
	"fmt"
	"time"
)

func main(){

	//1获取当前时间
	now:=time.Now() //2023-06-09 20:27:59.0141575 +0800 CST m=+0.004994201
	month:=time.Now().Month()// June
	day:=time.Now().Day()//9
	hour :=time.Now().Hour() //20
	minute:=time.Now().Minute() //32
	second := time.Now().Second()//32
	fmt.Println(now,month,day,hour,minute,second)

	//2locti

}