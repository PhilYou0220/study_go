package main

import "fmt"

//函数声明
func leijia(a int,b int) (sum int){
	sum = 0 //sum不用声明吗 a,b也没有声明，在函数头部就申明了
	for i:=a;i<=b;i++{
		sum = sum +i
	}
	return sum
}

func main() {
	ret :=leijia(1,100)
	fmt.Println(ret)
}
