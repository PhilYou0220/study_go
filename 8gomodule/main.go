

//import (
//	"fmt"
//	//"github.com/jinzhu/now"
//)

//go module
//启用 go module
//go env -w GO111MODULE=on
//或在编辑器settings里勾选go module
//到相应目录下执行 go mod init xxx
//初始化go module xxx一般是报名 会生成go.mod文件 执行这条命令时，当前目录不能存在go.mod文件（有删除）
//go mod tidy 导入外部包，外部包存在于GOPATH下 我的路径是E:\GoCode   E:\GoCode\pkg\mod\github.com\jinzhu
//go get "远程包" 用来下载包名
//go get "github.com/jinzhu/now"
//由于墙的原因 需要更换代理 go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/



package main
// go mod vendor 拉取包到本地环境；有点像 python的虚拟环境
import (
	"fmt"
	 "test/api1"
	"github.com/jinzhu/now" //使用第三方包时 需要导到本地
)


func main() {
	api1.RpcApi() //使用包名
	fmt.Println("1")
	fmt.Println(now.BeginningOfMinute()) //2022-11-20 21:08:00 +0800 CST
}
