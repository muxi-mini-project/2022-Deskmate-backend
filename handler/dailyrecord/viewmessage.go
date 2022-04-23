package dailyrecord //游览与当前同桌互相发送的消息

import (
	"Deskmate/handler"
	"Deskmate/model"
	"Deskmate/services/user"
	"strconv"
	"github.com/gin-gonic/gin"
)

// @Summary 浏览消息记录
// @Description 查看本次同桌消息记录，包含时间戳
// @Tags dailyrecord
// @Accept application/json
// @Produce application/json
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
		// c.JSON(401, gin.H{"message": "验证失败"})
		handler.SendResponse401(c, "Token Invalid.", nil) // 3.23
		return
	}

	
	DeskmateId, err := model.GetDeskmateId(id)
	if err != nil {
		/* c.JSON(401, gin.H{
			"message": "未查询到对应同桌关系id",
		}) */
		handler.SendResponse500(c, "未查询到对应同桌关系id", err) // 3.23
		return
	}

	//MessageInfor, err := model.GetMessage(DeskmateId)
	MessageInfor, err := model.GetReturnMessage(DeskmateId)
	/* if err != nil {
		// c.JSON(404, gin.H{"message": "获取失败"})
		handler.SendResponse500(c,"获取失败",err) // 3.23
		return
	} */ // 3.25 我将err判断去掉，看是否是数据库出错
	// c.JSON(200, MessgaeInfor)
	// LastMessage, err := model.GetLastMessgae(DeskmateId)
	if err != nil {
		handler.SendResponse500(c, "获取单条记录错误", err)
	}
	len := len(MessageInfor)
	for i := 0; i < len; i++ {
		if MessageInfor[i].UserId == id {
			MessageInfor[i].Name = "2"
		} else {
			MessageInfor[i].Name = "1"
		}
	}
	// handler.SendResponse(c, "success",LastMessage)
	days, err := model.GetDays(DeskmateId)
	//handler.SendResponse(c, "success", MessageInfor) // 3.23
	handler.SendResponse(c,strconv.Itoa(days), MessageInfor)
}
