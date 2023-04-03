package main

import "github.com/gin-gonic/gin"

func main(){
	//默认引擎
	r := gin.Default()
	//加载html
	r.LoadHTMLGlob("templates/*")

	name := "you"
	books :=[]string{"曾国藩的中年突围","红星照耀中国","陶庵梦忆"}
	//var stu01 = map[string]string{"name":"you","age":"18"}

	//声明一个map 键是string 值为空接口 啥子都能装
	stu_map := map[string]interface{}{
		"name":"phil","age":18,
	}

	type Person struct {
		Name string //大写外面才能读得到
		Age int64
	}
	person1 := Person{"you",18}
	//结构体的属性（相当于键）不用加 ""
	person2 := []Person{{Name:"you",Age:19},{Name:"san",Age:20}}
	r.GET("index", func(context *gin.Context) {
		context.HTML(200, "index.html",gin.H{
			"name":name,
			"books":books,
			"stu_map":stu_map,
			"Person":person1,
			"person2":person2,

		})
	})
	//启动
	r.Run("127.0.0.1:8081")

}


