package main

import "fmt"

func main() {


	//(1)获取地址 &
	var a = 10
	//fmt.Println(&a)
	//fmt.Printf("地址为：%p \n", &a)
	////(2)地址 声明 赋值
	var p *int // p为整型指针变量 这个根据你要取的地址的值的类型确定 如果是string 就要定义为*string
	p=&a
	fmt.Printf("p为%p \n",p)
	//
	////(3)取值操作 指向该地址存储的值
	//fmt.Println(*p)
	//fmt.Printf("值为%d \n",*p)
	//
	////可以修改*p的值从而改变a的值 两个的地址是一样的
	//*p = 100
	//fmt.Println(a)

	//练习
	//var v1 = 100 //  v1 int 100
	//var v2 =&v1  // v2 *int(v1存的是int类型 所以v2类型是*int) v1的地址
	//var v3 =&v2 // v3 **int(v2存的是*int(v1的地址)  所以v3的类型是**int) v2的地址
	//**v3 =200
	//fmt.Println(v1)

	p1 := 1
	p2 :=&p1
	*p2++ //先取值，再相加
	fmt.Println(p1)

}