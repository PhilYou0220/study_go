package main
import "github.com/gin-gonic/gin"


func main() {
	r:=gin.Default()
	//路由映射函数同一个路由，不同的方法执行不同的逻辑
	r.GET("/book", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "查询成功",
		})

	})
	r.POST("/book", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "新增成功",
		})

	})

	r.PUT("/book", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "修改成功",
		})

	})

	r.DELETE("/book", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "删除成功",
		})

	})
	// any请求方式都可以访问
	r.Any("/index", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "任何方式都可以访问",
		})

	})

	//所有路由都无法访问时，不管何种请求方式，走noroute时返回相应信息
	r.NoRoute(func(context *gin.Context) {
		context.JSON(404, gin.H{
			"msg": "404 not find",
		})

	})

}
