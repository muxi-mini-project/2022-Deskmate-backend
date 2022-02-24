package user //获取用户信息,比如点击某个标签跳出用户的信息

import (
	"Deskmate/model"
	"Deskmate/services/user"

	"github.com/gin-gonic/gin"
)

// @Summary 用户界面
// @Description "获取用户的基本信息"
// @Tags user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 "成功"
// @Failure 404 "获取失败"
// @Router /user [get]
func Userinfo(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "验证失败"})
	}
	Userinformation, err := model.GetUserInfo(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
	}
	c.JSON(200, Userinformation)
}
