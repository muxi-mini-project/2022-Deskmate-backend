package card //1.24 第一次设置名片，后面的是修改名片

import (
	"Deskmate/model"
	"Deskmate/services/user"
	"log"

	"github.com/gin-gonic/gin"
)

// @Summary 修改名片信息
// @Description "修改名片的昵称，标签，头像，用户宣言"
// @Tags card
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 "修改成功"
// @Failure 401 "验证失败"
// @Failure 400 "修改失败"
// @Router /card [PUT]
func Edit(c *gin.Context) {
	var card model.Card
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "验证失败"})
	}
	if err1 := c.BindJSON(&card); err1 != nil {
		c.JSON(400, gin.H{"message": "输入格式有误"})
		return
	}
	card.UserId = id
	if card.NickName == "" {
		c.JSON(400, gin.H{"message": "昵称不可为空！"})
		return
	}

	for _, char := range card.NickName {
		if string(char) == " " {
			c.JSON(400, gin.H{"message": "昵称不可为空格！"})
			return
		}
	}

	log.Println(card.UserId)

	if err2 := model.ChangeCardInfo(card); err2 != nil {
		c.JSON(400, gin.H{"message": "修改失败"})
		return
	}
	c.JSON(200, gin.H{"message": "修改成功"})
}
