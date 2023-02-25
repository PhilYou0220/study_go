package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

func main() {

	//序列化
	//var a string
	//a = "飞123" //这样是不能转的 json没这个格式

	var b = []int{1, 2, 3}
	content, _ := json.Marshal(b) //content 是个字节切片
	fmt.Println(content, string(content)) //[91 49 44 50 44 51 93] [1,2,3]
	fmt.Printf("%#v\n",content) //[]byte{0x5b, 0x31, 0x2c, 0x32, 0x2c, 0x33, 0x5d}
	fmt.Printf("%#v",string(content)) //"[1,2,3]"

	_:ioutil.WriteFile("test.json",content, 0666 ) //写入json content是[]byte

	c,_:=ioutil.ReadFile("test.json") //返回的是个[]byte
	fmt.Println(string(c)) //[1,2,3]
	var info = make([]int,100)
	 _= json.Unmarshal(c ,&info) //c[]byte,&info还需要指针类型的来接收
	 fmt.Println(info,reflect.TypeOf(info)) //[1 2 3] []int

}
