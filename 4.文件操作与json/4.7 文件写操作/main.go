package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeBytesOrStr(file *os.File) {
	str := "满江红666\n"
	//写入字节切片数据
	file.Write([]byte(str))
	//直接写入字符串数据
	file.WriteString("怒发冲冠,凭栏处、潇潇雨歇。")
}

func writeByBufio(file *os.File) {
	writer := bufio.NewWriter(file)
	//将数据先写入缓存，并不会到文件中
	writer.WriteString("大浪淘沙\n")
	// 必须flush将缓存中的内容写入文件
	writer.Flush()
}

func writeFile() {
	str := "怒发冲冠，凭栏处、潇潇雨歇。"
	err := ioutil.WriteFile("满江红", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func main() {

	file, err := os.OpenFile("满江红.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()

	// 写字节或者字符串
	writeBytesOrStr(file)
	// flush写
	writeByBufio(file)
	// 写文件
	writeFile()

}