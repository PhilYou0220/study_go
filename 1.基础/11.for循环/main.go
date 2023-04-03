package main

import "fmt"
//for range 可以遍历数组、切片、字符串、map 及通道（channel） 不能循环struct
func main() {

	//死循环
	//for  {
	//	fmt.Println("狗狗")
	//	time.Sleep(time.Second*1)
	//}

	// 只写条件可以
	//var a int = 10
	//for a>5{
	//	fmt.Println("很好")
	//	a--
	//}

	//变量;条件；变量赋值
	//for i := 1; i < 5; i++ {
	//	fmt.Println("很好")
	//	//i = i+1
	//}

	//break与continue
	//for i:=1;i<4;i++{
	//	for j:=1;j<4;j++{
	//
	//		if j==2{
	//			continue
	//		}
	//
	//		if i==3 {
	//			fmt.Println("i等于3")
	//			break
	//		}
	//		fmt.Println(i,j)
	//
	//	}
	//	fmt.Println(i)
	//}
	// go特性打标签
	// 111 211
f1:
	for i := 1; i < 3; i++ {
	f2:
		for j := 1; j < 3; j++ {
			if j == 2 {
				continue f1
			}
			fmt.Println(i, j)
			for k := 1; k < 3; k++ {

				if k == 2 {
					continue f2 //下次f2开始循环
				}
				fmt.Println(i, j,k)
			}
		}
	}

}
