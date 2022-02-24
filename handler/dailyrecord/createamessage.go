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
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 "发送成功"
// @Failure 401 "身份验证失败 重新登录"
// @Failure 400 "发送失败"
// @Router /dailyrecord/send [post]
func Createamessage(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "验证失败"})
	}
	//开始前先判断此时是否有同桌，如果没有，则告诉用户前往同桌广场寻找同桌
	result, err := model.GetUserStatus(id) //判断此时申请对象是否有同桌
	if err != nil {
		c.JSON(401, gin.H{
			"message": "未查询到申请对象是否有同桌.",
		})
		return //用return直接返回，不需要用If语句
	}
	if result == "无" {
		c.JSON(200, "你还没有同桌，请前往同桌广场寻找同桌!")
		return
	}

	DeskmateId, err := model.GetDeskmateId(id)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "未查询到对应同桌关系id",
		})
		return
	}

	//判断你今天是否发送过消息了，要求是一天只能发送一次信息
	judgement, err := model.GudgeSendRepeat(id, DeskmateId)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "未查询到相关信息",
		})
		return
	}
	if judgement == "false" {
		c.JSON(401, gin.H{
			"message": "你今天已经发送过信息！",
		})
		return
	}

	var message model.Message
	message.UserId = id                 //发该条消息学生的Id
	message.DailyrecordsId = DeskmateId //该条消息所属的同桌关系Id
	message.Time = time.Now()
	if err := c.BindJSON(&message); err != nil { //BindJSON,把客户端输入的数据(如postman的json格式数据)存入cardInfo结构体中
		c.JSON(400, gin.H{
			"message": "Lack Param Or Param Not Satisfiable.", //缺少参数或参数不可满足。
		})
		return
	}
	errs := model.DB.Create(&message)
	if errs != nil {
		c.JSON(400, "Fail.")
	}
	handler.SendResponse(c, "创建成功", nil)
}
