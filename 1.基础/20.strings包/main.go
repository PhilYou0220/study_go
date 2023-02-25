package main

import (
	"fmt"
	"strings"
)

func main() {
	//字符串 大小写转换
	s := "yoU are my love 张"
	//全转大写
	fmt.Println(strings.ToUpper(s)) //YOU ARE MY LOVE 张
	//全转小写
	fmt.Println(strings.ToLower(s)) //you are my love 张

	//是否以某个字符串开头
	fmt.Println(strings.HasPrefix(s, "yo")) //true
	//是否以某个字符串结尾
	fmt.Println(strings.HasSuffix(s, "张")) //true

	//是否包含某个字符
	fmt.Println(strings.Contains(s, "my")) //true

	//去除两端空格或其它字符 不论多少空格都能去 Trim 修剪
	s1 := "  you  "
	fmt.Println(strings.Trim(s1, " "))
	//直接去空格
	fmt.Println(strings.TrimSpace(s1))
	//去除左侧空格或其它字符
	fmt.Println(strings.TrimLeft(s1, " "))
	//去除右侧空格或其它字符
	fmt.Println(strings.TrimRight(s1, " "))

	//字符串索引Index查找
	s2 := "我叫张三"
	fmt.Println(strings.Index(s, "are"))
	fmt.Println(strings.Index(s2, "叫")) //3 中文亦如既往的三个字节代表一个汉字

	//分割 拼接
	s3 := "to be or not to be"
	s_s3 := strings.Split(s3, " ")
	fmt.Println(s_s3) //[to be or not to be] 生成了个切片 pyhton叫列表
	j_s3 := strings.Join(s_s3,",")
	fmt.Println(j_s3) //to,be,or,not,to,be 生成了个字符串
}