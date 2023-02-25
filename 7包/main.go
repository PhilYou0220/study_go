package main

//gopath需要此项目的绝对路径导包

import "study/7包/api"  // 使用gopath 需要执行
// GO111MODULE=off
//import ."study/7包/api"  //这样写就可以写api包名了


func main() {
	api.RpcApi() //使用包名
}
