package card //获取用户名片信息,比如点击某个标签跳出用户的名片信息

import (
	"Deskmate/model"
	"Deskmate/services/user"

	"github.com/gin-gonic/gin"
)

// @Summary 名片界面
// @Description "获取用户的名片信息"
// @Tags card
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} model.User
// @Failure 404 "获取失败"
// @Router /card [get]
func Cardinfo(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "验证失败"})
	}
	Userinformation, err := model.GetCardInfo(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
	}
	c.JSON(200, Userinformation)
}
