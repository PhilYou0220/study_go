package core

import "github.com/gin-gonic/gin"

func Book(context *gin.Context) {
	context.JSON(200, gin.H{
		"msg": "查询成功",
	})
}

func BookAdd(context *gin.Context) {
	context.JSON(200, gin.H{
		"msg": "新增成功",
	})
}

func BookEdit(context *gin.Context) {
	context.JSON(200, gin.H{
		"msg": "编辑成功",
	})
}

func BookDelete(context *gin.Context) {
	context.JSON(200, gin.H{
		"msg": "删除成功",
	})
}



