package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @title Deskmate
// @version 1.0.0
// @description 同桌小程序
// @termsOfService http://swagger.io/terrms/
// @contact.name qianren
// @contact.email 1911401642@qq.com
// @host localhost
// @BasePath api/v1/
// @Schemes http


func main() {
	// 1.创建带有默认中间件的路由； r :=gin.new()是创建没有中间件的路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}

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
