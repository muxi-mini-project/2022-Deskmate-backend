package card //1.24 第一次设置名片，后面的是修改名片

import (
	"Deskmate/model"
	"Deskmate/services/user"
	"Deskmate/handler"
	"log"

	"github.com/gin-gonic/gin"
)

// @Summary 修改名片信息
// @Description "修改名片的昵称，标签，头像，用户宣言"
// @Tags card
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param object body model.Cardinfor true "名片信息"
// @Success 200 "修改成功"
// @Failure 401 "验证失败"
// @Failure 400 "修改失败"
// @Router /card [PUT]
func Edit(c *gin.Context) {
	var card model.Card
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		// c.JSON(401, gin.H{"message": "验证失败"})
		handler.SendResponse401(c, "验证失败", err) // 3.22
		return
	}
	if err1 := c.BindJSON(&card); err1 != nil {
		// c.JSON(400, gin.H{"message": "输入格式有误"})
		handler.SendBadRequest(c, "Lack Param or Param Not Satisfiable.",err) // 3.22
		return
	}
	card.UserId = id
	if card.NickName == "" {
		// c.JSON(400, gin.H{"message": "昵称不可为空！"})
		handler.SendBadRequest(c, "昵称不可为空！", err) //3.22
		return
	}

	for _, char := range card.NickName {
		if string(char) == " " {
			// c.JSON(400, gin.H{"message": "昵称不可为空格！"})
			handler.SendBadRequest(c, "昵称不可为空！", err) //3.22
			return
		}
	}

	log.Println(card.UserId)

	if err2 := model.ChangeCardInfo(card); err2 != nil {
		// c.JSON(400, gin.H{"message": "修改失败"})
		handler.SendBadRequest(c, "修改失败！", err)
		return
	}
	// c.JSON(200, gin.H{"message": "修改成功"})
	handler.SendResponse(c, "修改成功!","")
}
