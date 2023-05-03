package main

import "fmt"

func main() {

	var  a string
	a  = "飞123"
	fmt.Println(len(a)) //6 len是数字节数，一个中文三个字节

	//三要素for循环也是循环字节
	for i:=0;i<len(a);i++{
		fmt.Println(a[i])
	}
	//233
	//163
	//158
	//49
	//50
	//51

	//想要循环字符使用for range
	//for range 可以遍历数组、切片、字符串、map 及通道（channel） 容器类型吧
	for i,v := range a{
		fmt.Println(i,v)
	}
	//0 39134
	//3 49
	//4 50
	//5 51
}
