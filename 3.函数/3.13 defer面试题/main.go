package main

import "fmt"
//defer 不能修改return的值 就是先不看 defer 先做 return

//rval = xxx
//defer_func
//ret rval

//defer x = 100
//x := 10
//return x  // rval=10.   x = 100, ret rval
func f1() int {
	i := 5
	defer func() {
		i++
	}()
	return i  //rval = 5,i=6 ret rval
}
func f2() *int {

	i := 5
	defer func() {
		i++
		fmt.Printf(":::%p\n", &i)
	}()
	fmt.Printf(":::%p\n", &i)
	return &i //rval= i的地址 ,i=6 ,ret rval
}

func f3() (result int) { //存在命名的返回值，会把return的值赋给result 就是说 result会取代rval 不存在率rval了
	defer func() {
		result++
	}()
	return 5 // result = 5;defer里result=6;  ret result(result替换了rval)
}

func f4() (result int) {
	defer func() {
		result++
	}()
	return result

	//  result=0; result=1 ;ret result
}

func f5() (r int) {
	t := 5
	defer func() {
		t = t + 1
	}()
	return t


	// r=t= 5 ; t=6; ret r = 5 (拷贝t的值5赋值给r)
}

func f6() (r int) {
	fmt.Println(&r)
	defer func(r int) { //不是闭包
		r = r + 1
		fmt.Println(&r)
	}(r)
	return 5

	// f6.r =5; func.r=1 ,ret f6.r
}

func f7() (r int) { //是闭包 r没有声明
	defer func(x int) {
		r = x + 1
	}(r)
	return 5  //f7.r =5;f7.r=1; ret f7,r
}

func main() {

	//println(f1()) //6 5
	//println(*f2()) //地址不一样 做错了 地址一样 6
	//println(f3()) //6 5 //又错了 命名返回值 有坑 6
	//println(f4()) //0  又做错 1
	//println(f5()) //6  又错 5
	println(f6())//地址一样 错 地址不一样 5
	//println(f7())//6 错 1

}
