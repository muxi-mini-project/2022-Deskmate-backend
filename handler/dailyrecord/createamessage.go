package dailyrecord

import (
	"Deskmate/handler"
	"Deskmate/model"
	"Deskmate/services/user"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary 发送消息
// @Description "发送一条消息"
// @Tags dailyrecord
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param object body model.Record true "每日打卡的内容"
// @Success 200 "发送成功"
// @Failure 401 "身份验证失败 重新登录"
// @Failure 400 "发送失败"
// @Failure 500 "服务器发生错误"
// @Router /dailyrecord/send [post]
func Createamessage(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		handler.SendResponse401(c,"Token Invalid.",nil) // 3.23
		return
	}

	//开始前先判断此时是否有同桌，如果没有，则告诉用户前往同桌广场寻找同桌
	result, err := model.GetUserStatus(id) //判断此时申请对象是否有同桌
	if err != nil {
		handler.SendResponse401(c,"未查询到申请对象是否有同桌.",err) // 3.23
		return //用return直接返回，不需要用If语句
	}
	
	if result == "0" {
		// c.JSON(200, "你还没有同桌，请前往同桌广场寻找同桌!")
		handler.SendResponse(c,"你还没有同桌，请前往同桌广场寻找同桌！",nil) // 3.23
		return
	}

	DeskmateId, err := model.GetDeskmateId(id)
	if err != nil {
		handler.SendResponse401(c,"未查询到对应同桌关系id",err) // 3.23
		return
	}

	//判断你今天是否发送过消息了，要求是一天只能发送一次信息
	judgement, err := model.GudgeSendRepeat(id, DeskmateId)
	if err != nil {
		handler.SendResponse401(c,judgement,err) // 3.23
		return
	}

	if judgement == "false" {
		handler.SendResponse401(c,"你今天已经发过信息了!",nil)
		return
	}

	var message model.Message
	message.UserId = id                 //发该条消息学生的Id
	message.DailyrecordsId = DeskmateId //该条消息所属的同桌关系Id
	message.Time = time.Now().Format("2006-01-02 15:04:00")
	if err := c.BindJSON(&message); err != nil { //BindJSON,把客户端输入的数据(如postman的json格式数据)存入cardInfo结构体中
		handler.SendBadRequest(c,"Lack Param Or Param Not Satisfiable.",nil) // 3.23
		return
	}

	if errs := model.DB.Create(&message).Error;errs != nil {
		// c.JSON(400, "Fail.")
		handler.SendResponse500(c,"Fail.",errs) // 3.23
		return
	}
	handler.SendResponse(c, "创建成功", nil)
	//handler.SendResponse(c, "创建成功", judgement) //4.12 测试judgement是否真的有值
}
