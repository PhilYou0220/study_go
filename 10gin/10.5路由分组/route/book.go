package route

import "github.com/gin-gonic/gin"
import . "route/core"
func InitBookRoute(r *gin.Engine){
	//路由分组 后面调用时候可以不写前缀了
	bookRoute:=r.Group("/books")
	bookRoute.GET("/", Book)
	bookRoute.POST("/", BookAdd)
	bookRoute.PUT("/", BookEdit)
	bookRoute.DELETE("/", BookDelete)
}
