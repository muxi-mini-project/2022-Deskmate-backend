package dailyrecord //游览与当前同桌互相发送的消息

import (
	"Deskmate/model"
	"Deskmate/services/user"

	"github.com/gin-gonic/gin"
)

// @Summary 浏览消息记录
// @Description 查看本次同桌消息记录，包含时间戳
// @Tags dailyrecord
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 "成功"
// @Failure 401  "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400  "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500  "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /dailyrecord/message [get]
func ViewMessage(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "验证失败"})
	}

	DeskmateId, err := model.GetDeskmateId(id)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "未查询到对应同桌关系id",
		})
		return
	}

	MessgaeInfor, err := model.GetMessage(DeskmateId)
	if err != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
	}
	c.JSON(200, MessgaeInfor)

}
