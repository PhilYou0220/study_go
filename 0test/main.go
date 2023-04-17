package main

import (
	"fmt"
	"time"
)

func main() {
	var arr = make([]int,5)
	fmt.Println(arr)

	fmt.Println(arr,len(arr),cap(arr)) //panic: runtime error: index out of range [0] with length 0
	fmt.Println(time.Now())

}
