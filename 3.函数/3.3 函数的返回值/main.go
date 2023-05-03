package main

import "fmt"

func test() {
	fmt.Println("test")
}

//return 语句中表达式的类型应与定义函数时指定的返回值类型必须一致。不然会报错
// 多个参数一个返回值  标注返回值类型
func test1(a, b int) int {
	c := a + b
	return c
}

//多个参数多个返回值 小括号标注其返回值
func test3(username, pwd string) (string,bool) {
	if username == "root" && pwd == "123" {
		return username, true
	} else {
		return username, false
	}

}
//面试题
func tets4(a,b int)(z int){
	fmt.Println(a,b)
	return //返回z 默认为0
}

func test5(a,b int)(sum,sub int){
	sum = a+b
	sub = a-b
	return //即使不写返回值规定了返回值 也会返回
}

func test6(a int){
	fmt.Println(a)
	return
}
func main() {
	//无返回值将报错
	//t :=test()
	//fmt.Println(t)

	re1 := test1(1, 2)
	fmt.Println(re1)

	user,b:=test3("root","123")
	if b{
		fmt.Println(user,b)
	}

	//定义了返回值z 应为0
	fmt.Println(tets4(1,2))

	// test5
	fmt.Println(test5(1,2))
	//test6
	//fmt.Println(test6(1))
}
