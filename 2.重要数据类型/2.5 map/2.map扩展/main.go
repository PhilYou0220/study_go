package main

import (
	"fmt"
	"strconv"
)

func main() {
	// key string value 字符串切片
	//data :=make(map[string][]string)
	//data["beijing"]=[]string{"朝阳","海定","昌平"}
	//data["shangdong"] = []string{"济南","青岛","威海"}
	//data["hebei"] = []string{"唐山","石家庄","保定"}
	////打印河北的第二个城市
	//fmt.Println(data["hebei"][1])
	////循环打印出每个省份和城市数量
	//for k,v := range data{
	//	fmt.Println(k,len(v))
	//}
	////添加一个新的省份和城市
	//data["sichuang"] = []string{"成都","遂宁","绵阳"}
	//fmt.Println(data)
	////删除北京的key-value
	//delete(data,"beijing")
	//fmt.Println(data)

	//case2 map嵌套map
	//info := map[int]map[string]string{1001: {"name": "yuan", "age": "23"}, 1002: {"name": "alvin", "age": "33"}}
	//// 打印学号为1002的学生的年龄
	//fmt.Println(info[1002]["age"])
	//// 循环打印每个学员的学号，姓名，年龄
	//for k,v := range info{
	//	fmt.Println(k)
	//	for a,b :=range v{
	//		fmt.Println(a,b)
	//	}
	// }
	//// 添加一个新的学员
	//info[1003] = map[string]string{"name":"fei","age":"26"}
	//fmt.Println(info)
	//// 删除1001的学生
	//delete(info,1001)
	//fmt.Println(info)

	// 案例3 切片嵌套map 一定要想到底层是个数组
	stus := []map[string]string{{"name": "yuan", "age": "23"}, {"name": "rain", "age": "22"}, {"age": "32", "name": "eric"}}
	// 打印第二个学生的姓名
	fmt.Println(stus[1]["name"])
	// 循环打印每一个学生的姓名和年龄
	for _,v :=range stus{
		fmt.Println(v["name"],v["age"])
	}
	// 添加一个新的学生map
	stus = append(stus,map[string]string{"name": "fei", "age": "26"})
	fmt.Println(stus) //[map[age:23 name:yuan] map[age:22 name:rain] map[age:32 name:eric] map[age:23=6 name:fei]]
	// 删除一个学生map
	stus = append(stus[:1],stus[2:]...)
	fmt.Println(stus) //[map[age:23 name:yuan] map[age:32 name:eric] map[age:23=6 name:fei]]
	//将姓名为yuan的学生的年龄自加一岁
	for _,b := range  stus{
		if b["name"]=="yuan"{
			age,_ := strconv.Atoi(b["age"])

			b["age"] = strconv.Itoa(age+1)
		}
	}
	fmt.Println(stus)



}
