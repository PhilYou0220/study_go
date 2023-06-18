package main

//用于对照probuf的嵌套

type Response struct {
	name string
	res []Result
}
type Result struct {
	data string
}