package main

import "fmt"


type Student struct {
	sid int
	name string
	age int
	course []string
}
func main() {
	//声明并赋值
	// (1) 方式1
	//s1 := Student{}
	var s1 = Student{}
	s1.sid = 1001
	s1.name = "yuan"
	fmt.Println(s1)
	// (2) 方式2：键值对赋值
	s2 := Student{sid: 1002, name: "rain", course: []string{"chinese", "math", "english"}}
	fmt.Println(s2)
	// (3) 方式3：多值赋值
	s3 := Student{1003, "alvin", 22, []string{"chinese", "math", "english"}}
	fmt.Println(s3)
}
