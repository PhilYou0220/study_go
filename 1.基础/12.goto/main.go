package main

import "fmt"

func main() {
	name := "fei"
	if name == "fei" {
		goto f2
	}
	name = "svip"
	if name == "svip" {
		goto f1

	}

f1:
	fmt.Println("很好")
f2:
	for i := 1; i < 3; i++ {
		fmt.Println("跳跃至此")
	}
}
