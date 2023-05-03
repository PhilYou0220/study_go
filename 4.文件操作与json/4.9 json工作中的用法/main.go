package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Stuno int `json:"stu_no"`
}

func main()  {
	//序列化  结构体往前端传参
	s1 := &student{Stuno: 18}
	data,err:=json.Marshal(s1)
	if err != nil{
		fmt.Println("出现问题了")
	}
	fmt.Println(data)
	fmt.Printf("json:%s \n",data)

	// 反序列化
	str := `{"stu_no":18}` //json串
	s2 := student{}
	err=json.Unmarshal([]byte(str),&s2)
	if err != nil{
		fmt.Println("解析失败！")
	}
	fmt.Println(s2)
	fmt.Printf("%#v",s2)


}