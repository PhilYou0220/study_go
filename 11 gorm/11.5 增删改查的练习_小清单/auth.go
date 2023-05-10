package main

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// 注册接口结构体
type RegParam struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//登录处理函数
func loginHandler(ctx *gin.Context) {
	//1获取参数
	var param RegParam
	if err := ctx.ShouldBind(&param); err != nil{
		ctx.JSON(http.StatusOK, Resp{
			Code: 1,
			Msg:  err.Error(),
		})
	}
	//2 逻辑处理 使用用户名和密码
	var user Account
	if err := db.Model(&Account{}).Where("name=? and pass_word=?", param.Name, md5Secret(param.Password)).Find(&user).Error;err != nil {
		//没查到
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusOK, Resp{
				Code: 1,
				Msg:  "用户名或密码错误",
			})
			return
		}

		//其余错误
		ctx.JSON(http.StatusOK, Resp{
			Code: 1,
			Msg:  err.Error(),
		})
		return

	}

	//登录成功，发送token
	token, err := GenToken(user.Uid, user.Name)
	if err!=nil{
		ctx.JSON(http.StatusOK, Resp{
			Code: 1,
			Msg:  "服务端异常，请稍后再试",
		})
		return
	}
	ctx.JSON(http.StatusOK, Resp{
		Code: 0,
		Msg: "登录成功",
		Data: token,
	})

}
func registerHandler(ctx *gin.Context) {
	//1获取参数
	var param RegParam
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusOK, Resp{
			Code: 1,
			Msg:  err.Error(),
		})
		return //参数不对不往下走
	}

	//2拿到参数后处理业务逻辑
	//查询是否有重复的用户名
	var user Account
	err := db.Model(&Account{}).Where("name=?", param.Name).First(&user).Error
	if err == nil {
		ctx.JSON(http.StatusOK, Resp{
			Code: 1,
			Msg:  "已存在此用户名",
		})
	}
	//其他错误
	if err != gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusOK, Resp{
			Code: 1,
			Msg:  "服务端异常，请稍后再试",
		})
	}
	//未找到此条数据就对了
	err = db.Create(&Account{
		Uid:      time.Now().Unix(),
		Name:     param.Name,
		PassWord: md5Secret(param.Password),
	}).Error

	if err != nil {
		ctx.JSON(http.StatusOK, Resp{
			Code: 1,
			Msg:  "服务端异常，请稍后再试",
		})
	}

	//注册成功
	ctx.JSON(http.StatusOK, Resp{
		Code: 0,
		Msg:  "注册成功",
	})
}

func md5Secret(pwd string) string {
	hash := md5.New()
	hash.Write([]byte(pwd))
	return hex.EncodeToString(hash.Sum([]byte(MySecret)))
}
