package route

import "github.com/gin-gonic/gin"

func InitRoute(r *gin.Engine){
	//路由分组
	InitBookRoute(r)

}
