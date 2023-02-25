package main

import "fmt"

func main() {
	age := 18
	switch age+1 {
	case 20:
		fmt.Println("非常不错")
	case 19:
		fmt.Println("good")
	default:
		fmt.Println("默认")

	}
}
