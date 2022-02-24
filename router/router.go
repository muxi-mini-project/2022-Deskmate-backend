package router //设置路径的文件

import (
	"Deskmate/handler/apply"
	"Deskmate/handler/card"
	"Deskmate/handler/dailyrecord"
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
		//登录用户，如果是首次登录将自动注册账号（一站式登录）
		g1.POST("", user.Login) //user.Login相当于一个API，每一个路由对应一个功能

		//查看个人信息，该信息是直接从学校读取的，不做修改
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

		//初始化用户头像
		g2.POST("/avatar", card.ModifyUserProfile)

		//修改用户头像

		g2.PUT("/avatar", card.ModifyUserProfile)
	}

	//square:同桌广场(显示名片信息和搜索tag)
	g3 := r.Group("/api/v1/square")
	{
		//名片信息流
		g3.GET("", square.ViewSquare)

		//搜索tag显示对应名片
		g3.GET("/tag", square.TagSearch)
	}

	g4 := r.Group("/api/v1/apply")
	{
		//创建同桌申请
		g4.POST("", apply.CreateApplication)

		//游览申请信息
		g4.GET("", apply.ViewApplication)

		//确认申请信息
		g4.PUT("", apply.ApplicationConfirm)
	}

	//dailyrecord:同桌每日打卡
	g5 := r.Group("/api/v1/dailyrecord")
	{
		//浏览同桌名片
		g5.GET("/card", dailyrecord.Cardinfo)

		//发送消息
		g5.POST("/send", dailyrecord.Createamessage)

		//浏览同桌消息
		g5.GET("/message", dailyrecord.ViewMessage)

		//提前终止同桌关系
		g5.PUT("/end", dailyrecord.BreakRelation)

		//更新打卡天数
		g5.PUT("/update", dailyrecord.Updatestatus)

	}
}
