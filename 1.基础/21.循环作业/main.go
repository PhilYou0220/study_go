package main

import "fmt"

func main() {

	//判断奇数 偶数
	//fmt.Println("请输入一个数字：")
	//var num int
	//_, err := fmt.Scanln(&num)
	//if err == nil {
	//	if num%2 ==0{
	//		fmt.Printf("%d偶数",num)
	//	}else   {
	//		fmt.Printf("%d奇数",num)
	//	}
	//}else {
	//	fmt.Println(err)
	//}

	//第二题
	//var  a int8 = 100
	//var b int16 = 200
	//fmt.Println(int16(a)+b)
	//
	////第三题 执行结果 //ls
	//cmd := ""
	//cmd += "l"
	//cmd += "s"
	//fmt.Println(cmd)

	//第四题
	//fmt.Println("请输入您的生日：如1990-3-16")
	//var s string
	//_, err := fmt.Scanln(&s)
	//if err != nil{
	//	fmt.Println(err)
	//}else{
	//	s1 := strings.Split(s,"-")
	//	fmt.Printf("您的生日是%s年-%s月-%s日",s1[0],s1[1],s1[2])
	//}

	//第5题
	//fmt.Println("请输入您的姓名如phil")
	//var name string
	//_, err := fmt.Scanln(&name)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	b := strings.ToUpper(name)
	//	if strings.HasPrefix(b,"A"){
	//		fmt.Println("true")
	//	}else {
	//		fmt.Println("false")
	//	}
	//
	//
	//}

	//第六题 flase
	//var name, age = "yuan", 22
	//fmt.Println(name == "yuan"  &&  age == 23)

	////第七题
	//var x =10
	//var y=5
	//x+=y
	//fmt.Println(x)

	//第八题
	//var a = 100
	//fmt.Println(strconv.Itoa(a))

	//第十四题 简单计算器 我写的太不优雅了
	//fmt.Println("请输入表达式如1+2：")
	//var s string
	//_,err:=fmt.Scanln(&s)
	//if err ==nil{
	//	if strings.Index(s,"+")>0 {
	//		a:= strings.Index(s,"+")
	//		b,_ :=strconv.Atoi(s[:a])
	//		c,_ :=strconv.Atoi(s[a+1:])
	//		fmt.Println(b+c)
	//	}else if strings.Index(s,"*")>0 {
	//		a:= strings.Index(s,"*")
	//		b,_ :=strconv.Atoi(s[:a])
	//		c,_ :=strconv.Atoi(s[a+1:])
	//		fmt.Println(b*c)
	//	}else if strings.Index(s,"/")>0 {
	//		a:= strings.Index(s,"/")
	//		b,_ :=strconv.Atoi(s[:a])
	//		c,_ :=strconv.Atoi(s[a+1:])
	//		fmt.Println(b/c)
	//	}else if strings.Index(s,"-")>0 {
	//		a:= strings.Index(s,"-")
	//		b,_ :=strconv.Atoi(s[:a])
	//		c,_ :=strconv.Atoi(s[a+1:])
	//		fmt.Println(b-c)
	//	}else {
	//		fmt.Println("格式错误",err)
	//	}
	//
	//} else{
	//	fmt.Println(err)
	//}

	//求1-10的阶乘 可以分析并得出结论 经典题
	re := 0
	for i:=1;i<=3;i++{
		tmp :=1
		for j:=i;j>1;j--{
			tmp = tmp*j
		}
	re = re+tmp
	}
	fmt.Println(re)

	}
