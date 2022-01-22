package router //设置路径的文件

import (
	"Deskmate/handler/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API router.")
	})
	//user:
	g1 := r.Group("/api/v1/user")
	{
		//登录
		g1.POST("", user.Login) //user.Login相当于一个API，每一个路由对应一个功能

		//
	}
}
