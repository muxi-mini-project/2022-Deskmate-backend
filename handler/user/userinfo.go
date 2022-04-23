package user //获取本人用户信息,比如点击某个标签跳出用户的信息

import (
	"Deskmate/handler"
	"Deskmate/model"

	"github.com/gin-gonic/gin"
)

// @Summary 用户界面
// @Description "获取用户的基本信息"
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param object body model.Id true "用户学号"
// @Success 200  {object} []model.User "搜索成功"
// @Failure 401 "身份验证失败"
// @Failure 404 "获取基本信息失败"
// @Router /user/infor [post]
func Userinfo(c *gin.Context) {
	/* token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		// c.JSON(401, gin.H{"message": "验证失败"})
		handler.SendResponse401(c,"身份验证失败",err) // 3.21
		return
	} */
	var id model.Id
	if err := c.BindJSON(&id); err != nil { //从客户端读取申请对象的学号
		handler.SendBadRequest(c, "Lack Param Or Param Not Satisfiable.", err)
		return
	}
	Userinformation, err := model.GetUserInfo(id.Id)
	
	if err != nil {
		// c.JSON(404, gin.H{"message": "获取失败"})
		handler.SendResponse404(c, "获取基本信息失败", err) //3.21
		return
	}
	// c.JSON(200, Userinformation)
	
	handler.SendResponse(c, "搜索成功", Userinformation) // 3.21
}
