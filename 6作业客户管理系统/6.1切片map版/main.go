package main

import (
	"fmt"
	"strings"
)

var userInfo = make([]map[string]string, 5)
var index int

func loginPage() {
	fmt.Println(`
-----------客户信息管理系统-------------
1.添加客户
2.查看客户
3.更新客户
4.删除客户
5.退出
-------------------------------------
`)

	choice()
}
func choice() {
	var number int
	fmt.Println("请输入选择【1-5】:")
	_, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println(err)

	}
	switch number {
	case 1:
		addUser()
		tips(addUser)
	case 2:
		seeUser()
		tips(seeUser)
	case 3:
		updateUser()
		tips(updateUser)
	case 4:
		deleteUser()
		tips(deleteUser)
	case 5:
		fmt.Println("程序已退出")
	}

}

func addUser() {
	var (
		name  string
		sex   string
		age   string
		email string
	)
	var user = make(map[string]string)
	fmt.Println("--------新增用户开始--------")
	fmt.Println("请输入新增用户姓名：")
	_, _ = fmt.Scan(&name)
	fmt.Println("请输入新增用户性别：")
	_, _ = fmt.Scan(&sex)
	fmt.Println("请输入新增用户年龄：")
	_, _ = fmt.Scan(&age)
	fmt.Println("请输入新增用户邮箱")
	_, _ = fmt.Scan(&email)

	user["name"] = name
	user["sex"] = sex
	user["age"] = age
	user["email"] = email
	userInfo[index] = user
	index++
	fmt.Println("--------新增用户结束--------")

}
func seeUser() {
	for k, v := range userInfo {
		if v != nil { //布尔类型的零值（初始值）为 false 数值类型的零值为 0 字符串类型的零值为空字符串 除了这桑默认值都是nil
			fmt.Printf("编号：%-d  姓名：%-s 性别：%-s 年龄：%-s 邮箱: %s \n", k+1, v["name"], v["sex"], v["age"], v["email"]) //-5左对齐5个
		}
	}

}

func updateUser() {
	fmt.Println("--------更新用户开始--------")
	var count int
	fmt.Println("请输入更新用户编号：")
	_,_=fmt.Scan(&count)
	if userInfo[count-1] !=nil {

		fmt.Println("请输入更新用户信息")
		var (
			name  string
			sex   string
			age   string
			email string
		)
		var user = make(map[string]string)
		fmt.Println("请输入新增用户姓名：")
		_, _ = fmt.Scan(&name)
		fmt.Println("请输入新增用户性别：")
		_, _ = fmt.Scan(&sex)
		fmt.Println("请输入新增用户年龄：")
		_, _ = fmt.Scan(&age)
		fmt.Println("请输入新增用户邮箱")
		_, _ = fmt.Scan(&email)
		user["name"] = name
		user["sex"] = sex
		user["age"] = age
		user["email"] = email
		userInfo[count-1] = user
		fmt.Println("--------更新用户结束--------")
	}else {
		fmt.Println("输入编号无数据！请重新输入")
		updateUser()
	}
}

func deleteUser() {
	fmt.Println("--------删除用户开始--------")
	var count int
	fmt.Println("请输入删除用户编号：")
	_,_=fmt.Scan(&count)

	//删除用户
	userInfo = append(userInfo[:count-1],userInfo[count:]...)
	fmt.Println("--------删除用户结束--------")
}

func tips(f func()) {
	fmt.Println("是否返回上一层【Y/N】:")
	var isNo string
	_, _ = fmt.Scan(&isNo)
	if strings.ToUpper(isNo) == "Y" {
		loginPage()
	} else if strings.ToUpper(isNo) == "N" {
		f()
		tips(f) //防止直接结束
	} else {
		fmt.Println("您的输入有误，请重新输入！")
		tips(f)
	}
}

func main() {

	loginPage()
	//index  := len(userInfo)
	//print(index)

}
