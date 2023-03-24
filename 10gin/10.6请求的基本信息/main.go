package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	User string `json:"name"` //传参传name 标记json就需要传 content-type json
	Pwd  int    `json:"pwd"`
}

func main() {
	r := gin.Default()

	//基本请求方式
	r.Any("/book", func(context *gin.Context) {

		//获取请求方式
		fmt.Println("method:", context.Request.Method) //method: GET
		//获取url 此方法是带url玩整路径如？a=1
		fmt.Println("method:", context.Request.URL) // method: /book?a=1
		//获取路由 此方法只获取路由 可以对路由做权限
		fmt.Println(context.FullPath()) //book
		//某个请求头
		fmt.Println("user-agent: ", context.GetHeader("user-agent")) //Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36
		//所有请求头
		fmt.Println("所有头：", context.Request.Header) //map[Accept-Encoding:[gzip, deflate, br]]
		//获取远程主机地址
		fmt.Println("", context.Request.RemoteAddr)   // 127.0.0.1:50429
		fmt.Println("RemoteIP: ", context.RemoteIP()) //127.0.0.1
		fmt.Println("ClientIP: ", context.ClientIP()) //127.0.0.1

		context.JSON(200, gin.H{
			"msg": "查询成功1",
		})

	})

	//获取get请求数据
	r.GET("/index", func(context *gin.Context) {

		//获取get请求方式数据 http://127.0.0.1:8081/index?kw=you
		fmt.Println("kw:", context.Query("kw")) //kw: you
		//设置默认值 取不到存在默认值 you
		fmt.Println(context.DefaultQuery("dkw:", "you")) //you
		context.String(200, "index")

		//取get数据如果没有提交则是false
		kw, ok := context.GetQuery("kw")
		if !ok {
			fmt.Println("参数不存在：", ok) //参数不存在： false
		}
		println(kw)

	})

	//获取post请求数据
	r.POST("/data", func(context *gin.Context) {
		//PostForm接收form表单的数据 如 Content-Type: application/x-www-form-urlencoded user=youfei&pwd=123 或Content-Type: multipart/form-data;{"pwd":"123","user":"youfei"}
		//如果直接发json数据则不可以 Content-Type: application/json;
		user := context.PostForm("user")
		pwd := context.PostForm("pwd")
		//设置默认值 取不到存在默认值
		fmt.Println(context.DefaultPostForm("user", "'fei"))
		//取get数据如果没有提交则是false
		fmt.Println(context.GetPostForm("user"))
		//获取同一个键的多个值 返回切片
		fmt.Println(context.PostFormArray("user")) //[youfei zhang]

		context.JSON(200, gin.H{
			"user": user,
			"pwd":  pwd,
		})

	})

	//shouldbind函数 根据content-type解析 需要结构体 常用
	r.POST("/shouldBind", func(context *gin.Context) {

		user := User{}
		//关于json的反序列化--映射到结构体对象上去 shouldbind根据结构体里字段的标签来处理是何种类型的content-type 如果是form表单，那么就要标记为 `form`
		context.ShouldBind(&user)
		context.JSON(200, gin.H{
			"data": user,
		})

	})

	r.Run("127.0.0.1:8081")

}
