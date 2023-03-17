package main

import (
	"fmt"
	"net"
)
// 模拟符合http协议的服务端返回值
//浏览器如何解析要看content-type:text/html 解析成html,text/plain 解析成字符串,application/json 解析成json
//服务器的content-type很重要
func main() {
	//服务端
	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		//panic("创建失败!") //panic 是一个内置函数，用于在程序发生无法处理的错误时，引发一个运行时错误
		fmt.Printf("%s:",err)
	}
	defer listener.Close()
	for true {
		//获取客户端套接字对象，三次握手
		fmt.Println("server is waiting...")
		conn, err := listener.Accept()
		if err != nil {
			panic("连接失败！")
		}

		//接受客户端消息
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		fmt.Printf("data:%s\n", string(data[:n]))

		//返回信息 需要符合响应协议格式才行 因为是windows所以换行符是\r\n linux上就是\r
		//res := "hello world"
		//res := "HTTP/1.1 200 OK\r\ncontent-type:text/plain \r\n\r\n hello world"
		res := "HTTP/1.1 200 OK\r\ncontent-type:text/html \r\n\r\n <h1>hello world</h1>"
		_, err = conn.Write([]byte(res))
		if err != nil {
			fmt.Println("写入失败！", err)
			continue
		}
		conn.Close()

	}
}
