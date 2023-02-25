package main

import (
	"fmt"
	"strconv"
)

func main() {
	name := "游飞"
	fmt.Println(name[0],strconv.FormatInt(int64(name[0]),2))
	fmt.Println(name[1],strconv.FormatInt(int64(name[0]),2))
	fmt.Println(name[2],strconv.FormatInt(int64(name[0]),2))
	fmt.Println(name[3],strconv.FormatInt(int64(name[0]),2))
	fmt.Println(name[4],strconv.FormatInt(int64(name[0]),2))
	fmt.Println(name[5],strconv.FormatInt(int64(name[0]),2))
}