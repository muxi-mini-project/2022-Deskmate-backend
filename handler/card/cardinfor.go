package card //获取用户名片信息,比如点击某个标签跳出用户的名片信息

import (
	"Deskmate/model"
	"Deskmate/handler"
	"Deskmate/services/user"

	"github.com/gin-gonic/gin"
)

// @Summary 他人名片
// @Description "获取他人的名片信息"
// @Tags card
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param object body model.Id true "要查看用户的学号"
// @Success 200 {object} model.Card
// @Failure 401 "身份验证失败"
// @Failure 404 "获取失败"
// @Router /card/infor [post]
func Cardinfor(c *gin.Context) {
	token := c.Request.Header.Get("token")
	_, err := user.VerifyToken(token)
	if err != nil {
		// c.JSON(401, gin.H{"message": "验证失败"})
		handler.SendResponse401(c,"Token Invalid.",err) // 3.23
		return
	}

	var user_id model.Id
	if err := c.BindJSON(&user_id); err != nil { //从客户端读取申请对象的学号
		handler.SendBadRequest(c,"Lack Param Or Param Not Satisfiable.",err)
		return
	}

	Userinformation, err := model.GetCardInfo(user_id.Id)
	
	if err != nil {
		// c.JSON(404, gin.H{"message": "获取失败"})
		handler.SendResponse404(c,"获取失败",err) // 3.23
		return
	}

	// c.JSON(200, Userinformation)
	handler.SendResponse(c,"获取成功",Userinformation)
}
