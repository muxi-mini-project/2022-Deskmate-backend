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
// @host 119.3.2.168:4016
// @BasePath /api/v1/
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

func init() {
	port := flag.String("port", "4016", "本地监听的端口")
	// platform := flag.String("platform", "gitee", "平台名称，支持gitee/github")
	platform := flag.String("platform", "github", "平台名称，支持gitee/github")
	// token := flag.String("token", "411e5146ac55a87ff07b32e77e6d93e7", "Gitee/Github 的用户授权码")
	token := flag.String("token", "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", "Gitee/Github 的用户授权码")
	// owner := flag.String("owner", "ripples-of-year", "仓库所属空间地址(企业、组织或个人的地址path)")
	owner := flag.String("owner", "SUIYUELIANYI", "仓库所属空间地址(企业、组织或个人的地址path)")
	// repo := flag.String("repo", "gitee-picture-bed", "仓库路径(path)")
	repo := flag.String("repo", "DeskmateImage", "仓库路径(path)")
	// path := flag.String("path", "", "文件的路径")
	// branch := flag.String("branch", "master", "分支")
	branch := flag.String("branch", "main", "分支")
	flag.Parse()
	flag_handle.PORT = *port
	flag_handle.OWNER = *owner
	flag_handle.REPO = *repo
	// flag_handle.PATH = *path
	flag_handle.TOKEN = *token
	flag_handle.PLATFORM = *platform
	flag_handle.BRANCH = *branch
	if flag_handle.TOKEN == "" {
		panic("token 必须！")
	}
}
