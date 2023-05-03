package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readByByte(file *os.File )  {
	data := make([]byte,9,9)
	file.Read(data)
	fmt.Println(string(data))
}

func readByLines(file *os.File) {
	reader := bufio.NewReader(file)
	for {lineContent,err:=reader.ReadString('\n')
		fmt.Println(lineContent)
		if err == io.EOF{ //读到最后一行就是EOF
			break
		}
	}

}
func main() {
	//(1)读文件
	file,err:=os.Open("满江红") // file的类型是*File
	if err != nil {
		fmt.Println("打开文件失败！")
	}else {
		fmt.Println(file)
	}
	//1.1 按字节读
	//readByByte(file)

	//1.2按行读
	//readByLines(file)

	//1.3 读取整个文件
	content,_ :=ioutil.ReadFile("满江红") // 入参是string 上面两种方法的入参都是*file
	fmt.Println(string(content))
}
