package card //1.24 第一次设置名片，后面都是修改名片

import (
	"Deskmate/handler"
	"Deskmate/model"
	"Deskmate/services/user"
	"log"

	"github.com/gin-gonic/gin"
)

// @Summary 设置名片
// @Description "设置名片的昵称，标签，头像，用户宣言"
// @Tags card
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 "设置成功"
// @Failure 401 "身份验证失败 重新登录"
// @Failure 400 "设置失败"
// @Router /card [post]
func Create(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token) //这里用的是写在service里的解析token函数,并获取用户id即学号
	log.Println(id)

	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid"})
		return
	}
	var cardInfo model.Card
	cardInfo.UserId = id //这里直接从token里面调用学号放入名片，不需要在客户端手动输入
	cardInfo.Status = "无"
	if err := c.BindJSON(&cardInfo); err != nil { //BindJSON,把客户端输入的数据(如postman的json格式数据)存入cardInfo结构体中
		c.JSON(400, gin.H{
			"message": "Lack Param Or Param Not Satisfiable.", //缺少参数或参数不可满足。
		})
		return
	}
	result := model.DB.Create(&cardInfo)
	//result := model.ChangeCardInfo(cardInfo)
	if result != nil {
		c.JSON(400, "Fail.")
	}
	handler.SendResponse(c, "创建成功", nil)
}
