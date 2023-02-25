package main

import "fmt"

type Person struct{
	string // 相当于 string string
	int // 相当于 int int

}




func main() {
	p1 := Person{"fei",18}
	fmt.Print(p1)

}
