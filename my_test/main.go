package main

import "fmt"

type Todo struct {
Title  string
Status bool
}

func main()  {
	var todo Todo
	todo =Todo{}
	fmt.Println(todo)
}




