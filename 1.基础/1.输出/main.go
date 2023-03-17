package main

import "fmt"

func main() {
	//print("内置print")
	//println("内置println")
	fmt.Print("print")
	fmt.Println("println")
	fmt.Printf("老汉开着%s，去接alex这个%d货。\n", "兰博基尼", 2)

	//格式化字符串
	a := 1
	b := "张三"
	fmt.Printf("他是%s,%d岁了", b, a) //他是张三,1岁了

	//有个返回值 可以进行其它操作
	s := fmt.Sprintf("他是%s,%d岁了", b, a)
	fmt.Println(s)

	//多行字符串
	s1 := `第一行
			第二行
			第三行
			`
	fmt.Println(s1)
	//%v	值的默认格式表示
	//%+v	类似%v，但输出结构体时会添加字段名
	//%#v	值的Go语法表示
	test1 := map[string]string{"name":"you","age":"18"}
	fmt.Printf("%v\n",test1) //map[age:18 name:you]
	fmt.Printf("%+v\n",test1) //map[age:18 name:you]
	fmt.Printf("%#v\n",test1) //map[string]string{"age":"18", "name":"you"}
}
