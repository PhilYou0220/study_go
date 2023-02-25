package main

import (
	"encoding/json"
	"fmt"
)



type Person struct {
	Name string `json:"name"` //这里加标签变为小写
	Age  int `json:"-"` //这会不显示
	Addr //匿名字段 结构体嵌套 相当于 Addr Addr
}

type Addr struct {
	Province string
	City     string 
	County   string
}

func (a *Addr) history() {
	fmt.Println("这是一座历史悠久的城市")
}
func main() {
	//特别注意go语言 字段小写是private ,大写是public

	p1 := Person{Name:"fei", Age: 18, Addr:Addr{Province: "四川", City: "遂宁", County: "射洪"}}
	//p1 := Person{name:"fei", age: 18, Addr:Addr{province: "sichuan", city: "suining", county: "shehong"}}

	
	//序列化
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))  //{"Name":"fei","Age":18,"Province":"四川","City":"遂宁","County":"射洪"}
	//序列化字段变为
	
	
	
	//反序列化
	var Per Person
	_=json.Unmarshal(p1Json,&Per)
	fmt.Println(Per)
}
