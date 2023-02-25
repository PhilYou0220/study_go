package main

import "fmt"

func main() {
	a := make(map[string]string)
	a["name"] = "san"
	a["age"] = "18"
	value,exist:= a["1"]
	if exist{
		fmt.Println(value,exist)
	}else {
		fmt.Println(value,exist)
	}

}
