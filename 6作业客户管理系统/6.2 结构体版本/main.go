package main

import (
	"fmt"
	"strings"
)

//面向对象
type UserInfo struct {
	Cid    int
	Name   string
	Age    int
	Gender string
	Email  string
}

//建立一个user的服务的结构体 基于此服务结构体发出增删改查动作
type UserService struct {
	Users []UserInfo //结构体切片
	Uid   int
}

func (us *UserService) loginPage() {
	fmt.Println(`
-----------客户信息管理系统-------------
	1.添加客户
	2.查看客户
	3.更新客户
	4.删除客户
	5.退出
-------------------------------------
`)

	us.choice()
}
func (us *UserService) choice() {
	var number int
	fmt.Println("请输入选择【1-5】:")
	_, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println(err)

	}
	switch number {
	case 1:
		us.addUser()
		us.tips(us.addUser)
	case 2:
		us.seeUser()
		us.tips(us.seeUser)
	case 3:
		us.updateUser()
		us.tips(us.updateUser)
	case 4:
		us.deleteUser()
		us.tips(us.deleteUser)
	case 5:
		fmt.Println("程序已退出")
	}
}

func (us *UserService) addUser() {
	var (
		name  string
		sex   string
		age   int
		email string
	)
	fmt.Println("--------新增用户开始--------")
	fmt.Println("请输入新增用户姓名：")
	_, _ = fmt.Scan(&name)
	fmt.Println("请输入新增用户性别：")
	_, _ = fmt.Scan(&sex)
	fmt.Println("请输入新增用户年龄：")
	_, _ = fmt.Scan(&age)
	fmt.Println("请输入新增用户邮箱")
	_, _ = fmt.Scan(&email)
	us.Uid++
	var newUserInfo = UserInfo{
		Cid:    us.Uid,
		Name:   name,
		Age:    age,
		Gender: sex,
		Email:  email,
	}
	//索引+1

	us.Users = append(us.Users, newUserInfo)
	fmt.Println("--------新增用户结束--------")

}
func (us *UserService) seeUser() {
	for _, v := range us.Users {
		fmt.Printf("编号：%-d  姓名：%-s 性别：%-s 年龄：%-d 邮箱: %s \n", v.Cid, v.Name, v.Gender, v.Age, v.Email) //-5左对齐5个

	}

}

func (us *UserService) updateUser() {
	fmt.Println("--------更新用户开始--------")
	var count int
	fmt.Println("请输入更新用户编号：")
	_, _ = fmt.Scan(&count)
	var updateIndex = -1
	for index, user := range us.Users {
		if user.Cid == count {
			updateIndex = index
		}
	}
	if updateIndex == -1 {
		fmt.Println("该编号不存在！")
		us.updateUser()
	}

	fmt.Println("请输入更新用户信息")
	var (
		name  string
		sex   string
		age   int
		email string
	)

	fmt.Println("请输入新增用户姓名：")
	_, _ = fmt.Scan(&name)
	fmt.Println("请输入新增用户性别：")
	_, _ = fmt.Scan(&sex)
	fmt.Println("请输入新增用户年龄：")
	_, _ = fmt.Scan(&age)
	fmt.Println("请输入新增用户邮箱")
	_, _ = fmt.Scan(&email)
	var newUserInfo = UserInfo{
		Cid:    count,
		Name:   name,
		Age:    age,
		Gender: sex,
		Email:  email,
	}
	us.Users[updateIndex] = newUserInfo
	fmt.Println("--------更新用户结束--------")

}

func (us *UserService) deleteUser() {
	fmt.Println("--------删除用户开始--------")
	var count int
	fmt.Println("请输入删除用户编号：")
	_, _ = fmt.Scan(&count)

	//删除用户
	us.Users = append(us.Users[:count-1], us.Users[count:]...)
	fmt.Println("--------删除用户结束--------")
}

func (us *UserService) tips(f func()) {
	fmt.Println("是否返回上一层【Y/N】:")
	var isNo string
	_, _ = fmt.Scan(&isNo)
	if strings.ToUpper(isNo) == "Y" {
		us.loginPage()
	} else if strings.ToUpper(isNo) == "N" {
		f()
		us.tips(f) //防止直接结束
	} else {
		fmt.Println("您的输入有误，请重新输入！")
		us.tips(f)
	}
}

func main() {
	//实例化
	us := UserService{Uid: 0, Users: []UserInfo{}}
	us.loginPage()

}
