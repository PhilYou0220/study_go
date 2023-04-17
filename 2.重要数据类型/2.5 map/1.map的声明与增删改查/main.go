package main

import "fmt"
//相当于python中的字典 map_name =map[key_type]value_type
func main() {
	//map的声明并初始化 map没有索引 无序的
	var stu01 = map[string]string{"name":"you","age":"18"} //
	// key查询
	fmt.Println(stu01["name"]) //you
	// 增加
	stu01["gender"] = "male"
	fmt.Println(stu01) //map[age:18 gender:male name:you]
	//修改
	stu01["gender"] = "female"
	fmt.Println(stu01) //map[age:18 gender:female name:you]
	// 删除(没找到键也不会报错)
	delete(stu01,"gender")
	fmt.Println(stu01) //map[age:18 name:you]

	//声明方式二 （只要内存中存了地址的都是引用类型）
	//new函数返回指针地址 make 返回引用类型 如map返回map 切片返回切片
	//stu02 := make(map[string]string)
	//stu02["name"] = "zhang"
	//fmt.Println(stu02) //map[name:zhang]

	//使用空接口可以使值不规定类型
	stu03 := make(map[string]interface{})
	stu03["age"]=11
	fmt.Println(stu03) //map[age:11]
	stu03["name"]="you"
	//循环遍历容器
	for k,v:=range stu03{
		fmt.Println(k,v)
	}
}
