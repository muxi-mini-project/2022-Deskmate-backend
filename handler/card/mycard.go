package card //获取用户名片信息,比如点击某个标签跳出用户的名片信息

import (
	"Deskmate/model"
	"Deskmate/handler"
	"Deskmate/services/user"

	"github.com/gin-gonic/gin"
)

// @Summary 我的名片
// @Description "获取自己的名片信息"
// @Tags card
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} model.Card
// @Failure 401 "身份验证失败"
// @Failure 404 "获取失败"
// @Router /card [GET]
func MyCardinfor(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		handler.SendResponse401(c,"Token Invalid.",err) 
		return
	}

	Userinformation, err := model.GetCardInfo(id)
	if err != nil {
		handler.SendResponse404(c,"获取失败",err) 
		return
	}

	handler.SendResponse(c,"获取成功",Userinformation)
}
