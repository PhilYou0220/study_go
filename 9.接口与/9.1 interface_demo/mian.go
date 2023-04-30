package main

import "fmt"

/*接口是一种类型 注重于我能干什么 所以一般加上er */
type student struct {}

func (s student)dream(){
	fmt.Println("学生在做梦")
}

//func makeDream(s student){
//	s.dream()
//}




type teacher struct {}

func (t teacher)dream(){
	fmt.Println("老师在做梦")
}

type  dreamer interface { //定义dreamer接口类型 能够调用dream（）函数 注重于我能干什么-->做梦
	dream() //注意函数接收值和返回值必须一致才行
}

func makeDream(x dreamer){ //此函数只关心传进来的参数能调用dream方法 你传进来是啥就是啥 能掉dream()就行
	x.dream()
}

func main(){
	var s student  //声明一个student 类型的变量 s
	var t teacher
	makeDream(s)
	makeDream(t)
}