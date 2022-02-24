package main

import (
	"Deskmate/config"
	"Deskmate/model"
	"Deskmate/router"
	"Deskmate/services/flag_handle"
	"flag"

	//"Deskmate/config"//参照hjj的
	"fmt"
	//"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/spf13/viper"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Deskmat
// @version 1.0.0
// @description 同桌小程序
// @termsOfService http://swagger.io/terrms/
// @contact.name qianren
// @contact.email 1911401642@qq.com
// @host localhost
// @BasePath api/v1/
// @Schemes http

//下面这个是根据wz和hjj的登录方式
func main() {
	r := gin.Default() //创建带有默认中间件的路由
	config.ConfigInit()
	//注意大小写规范
	model.DB = model.Initdb()

	router.Router(r)
	if err := r.Run(":4016"); err != nil {
		fmt.Println(err)
	}
}

//var err error

//这个是根据wyx写的登录
//func main() {
// 1.创建带有默认中间件的路由； r :=gin.new()是创建没有中间件的路由
/*r := gin.Default()
// 2.绑定路由规则，执行的函数
// gin.Context，封装了request和response
r.GET("/", func(c *gin.Context) {
	c.String(http.StatusOK, "hello World!")
})
// 3.监听端口，默认在8080
// Run("里面不指定端口号默认为8080")

model.DB = model.Initdb()
defer model.DB.Close()

r.Run(":8000")*/

//本地连接数据库
/*dsn := "root:123456@tcp(127.0.0.1:3306)/deskmate?charset=utf8mb4&parseTime=ture&loc=Local"
model.DB, err = gorm.Open("mysql", dsn)
if err != nil {
	fmt.Println("数据库连接失败!")
	panic(err)
}
//model.DB.AutoMigrate(model.DB)
r := gin.Default()
router.Router(r)
r.Run(":8080")
defer model.DB.Close()*/
//}

//下面这个只是一个示例
/*func main() {
	router := gin.Default()

	//路由必须为/user/{参数} 不可以为/user/或/user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name") //通过Param获取
		c.String(http.StatusOK, "Hello %s", name)
	})

	// *参数名称 表示该参数为可选参数, /user/john/ 和 /user/john/send 都可以
	// 如果不存在路由 /user/{:name}, 它将会重定向至 /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.Run(":8080")
}*/

func init() {
	port := flag.String("port", "4016", "本地监听的端口")
	platform := flag.String("platform", "gitee", "平台名称，支持gitee/github")
	token := flag.String("token", "411e5146ac55a87ff07b32e77e6d93e7", "Gitee/Github 的用户授权码")
	owner := flag.String("owner", "ripples-of-year", "仓库所属空间地址(企业、组织或个人的地址path)")
	repo := flag.String("repo", "gitee-picture-bed", "仓库路径(path)")
	path := flag.String("path", "", "文件的路径")
	branch := flag.String("branch", "master", "分支")
	flag.Parse()
	flag_handle.PORT = *port
	flag_handle.OWNER = *owner
	flag_handle.REPO = *repo
	flag_handle.PATH = *path
	flag_handle.TOKEN = *token
	flag_handle.PLATFORM = *platform
	flag_handle.BRANCH = *branch
	if flag_handle.TOKEN == "" {
		panic("token 必须！")
	}
}
