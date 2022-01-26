package router //设置路径的文件

import (
	"Deskmate/handler/card"
	"Deskmate/handler/square"
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
		//登录用户
		g1.POST("", user.Login) //user.Login相当于一个API，每一个路由对应一个功能

		//获取个人信息
		g1.GET("", user.Userinfo)
	}
	//card:
	g2 := r.Group("/api/v1/card")
	{
		//设置新的名片
		g2.POST("", card.Create) //post为创建

		//修改名片
		g2.PUT("", card.Edit) //put为修改

		//获取名片信息
		g2.GET("", card.Cardinfo)
	}
	//square:同桌广场(显示名片信息和搜索tag)
	g3 := r.Group("/api/v1/square")
	{
		//名片信息流
		g3.GET("", square.ViewSquare)

		//搜索tag显示对应名片
		g3.GET("tag", square.TagSearch)
	}
}
