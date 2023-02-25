package main

func main() {
	//defer语句是go语言提供的一种用于注册延迟调用的机制，是go语言中一种很有用的特性.就是最后执行
	//case1 defer最后执行 1 3 2
	//fmt.Println("fei01")
	//defer fmt.Println("fei02")
	//fmt.Println("fei03")

	//case2 多个defer  先注册最后执行先进后出  1 3 5 4 2
	//fmt.Println("fei01")
	//defer fmt.Println("fei02")
	//fmt.Println("fei03")
	//defer fmt.Println("fei04")
	//fmt.Println("fei05")
}
