package main

func main(){
	//HTTP(Hyper Text Transfer Protocol) 超文本传输协议。
	//1 基于TCP/IP协议之上的应用层协议。
	//2 基于请求--响应模式。 客户端发起 服务端才响应
	//3 短连接和长连接 HTTP 1.0 一次性连接，任务结束就中断；HTTP 1.1 默认使用长连接 客户端和服务端首部的Connection都设置为keep-alive,才支持。其实是多个http请求服用同一个TCP连接，节省开支。
	//4 无状态 没有信息记录


	//HTTP请求协议
	//1 请求首行
	//1.1 请求方式 get post
	//1.2 请求路由 /api/vi
	//1.3 HTTP版本 1.1/1.0
	//2.请求头  content_type:application/json
	//浏览器版本、时间、请求体数据格式（服务器好解析content_type）等
	//3 请求体
	//get在地址栏，post等请求数据才在这里


	//HTTP响应协议
	//1.响应首行
	//HTTP版本 1.1/1.0,状态码，
	//2响应头
	//浏览器版本、时间、请求体数据格式（客户端解析content_type）等
	//3响应体
	//{"data":1}


}

