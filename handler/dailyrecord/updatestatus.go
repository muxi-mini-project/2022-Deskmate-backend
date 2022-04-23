package dailyrecord //用于更新每天的同桌状态，如果不符合则立即结束，如果符合则天数加一，如果14天已满则可以再次发送申请是否继续成为同桌

import (
	"Deskmate/handler"
	"Deskmate/model"
	"Deskmate/services/user"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary 更新打卡
// @Description "每日更新打卡天数"
// @Tags dailyrecord
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 "更新成功"
// @Failure 401 "身份验证失败 重新登录"
// @Failure 400 "更新失败，同桌关系已经解除"
// @Failure 500 "失败"
// @Router /dailyrecord/update [put]
func Updatestatus(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		handler.SendResponse401(c, "Token Invalid.", nil)
		return
	}

	result, err := model.GetUserStatus(id) //判断自己是否有同桌
	if err != nil {
		handler.SendResponse500(c, "查询自己状态失败", err)
		return
	}

	if result == "0" {
		handler.SendResponse(c, "你还没有同桌，请前往同桌广场寻找同桌！", nil)
		return
	}

	DeskmateId, err := model.GetDeskmateId(id)
	if err != nil {
		handler.SendResponse500(c, "未查询到对应同桌关系id", err)
		return
	}

	interval, err := model.GudgeUpdateRepeat(DeskmateId) //interval是时间间隔，1表示今天未更新，0表示今天已经更新

	if interval == "0" {
		handler.SendResponse(c, "今天已经更新过了", err)
		return
	}

	sid1, err := model.GetPartner(DeskmateId, id) //sid1存储对方的id
	if err != nil {
		handler.SendResponse500(c, "未查询到相关成员！", err) // 3.23
		return
	}

	sid2 := id //sid2存储我的id

	days, err := model.GetDays(DeskmateId)
	if err != nil {
		/* c.JSON(500, gin.H{
			"message": "未查询到打卡天数",
			"error":   err,
		}) */
		handler.SendResponse500(c, "未查询到打卡天数", err) // 3.23
		return
	}

	//判断两人是否都在昨天打卡
	judgement1, err := model.GudgeYesterday(sid1, DeskmateId)
	if err != nil {
		handler.SendResponse500(c, "未查询到你同桌昨天的打卡记录", err) // 3.23
		return
	}

	judgement2, err := model.GudgeYesterday(sid2, DeskmateId)
	if err != nil {
		handler.SendResponse500(c, "未查询到你昨天的打卡记录", err) // 3.23
		return
	}

	if judgement1 == "false" || judgement2 == "false" {
		err = model.ChangeDeskmateStatus(DeskmateId)
		if err != nil {
			handler.SendResponse500(c, "未成功解除同桌关系", err) // 3.23
			return
		}

		err = model.ChangeStatusAgain(sid1)
		if err != nil {
			handler.SendResponse500(c, "未成功解除同桌(对方的状态未改变)", err) // 3.23
			return
		}

		err = model.ChangeStatusAgain(sid2)
		if err != nil {
			handler.SendResponse500(c, "未成功解除同桌关系(你的状态未更新)", err) // 3.23
			return
		}
		handler.SendBadRequest(c, "很遗憾，有人未按照要求及时打卡，本次同桌打卡终止!", nil) //3.23
		return
	}

	if judgement1 == "ture" && judgement2 == "ture" {
		days++
		err = model.UpdateDays(DeskmateId, days)
		if err != nil {
			handler.SendResponse500(c, "更新打卡错误", err) // 3.23
			return
		}
		var update model.Update
		update.DailyrecordsId = DeskmateId
		update.Time = time.Now().Format("2006-01-02 15:04:00")
		if errs := model.DB.Create(&update).Error; errs != nil {
			handler.SendResponse500(c, "Create update record fail.", errs)
			return
		}
		//判断你今天是否发送过消息了，要求是一天只能发送一次信息
		judgement, err := model.GudgeSendRepeat(id, DeskmateId)
		var a string //a为1表示今天还没打卡，a为0表示今天已经打卡
		if err != nil {
			handler.SendResponse401(c, judgement, err) // 3.23
			return
		}

		if judgement == "false" {
			a = "0"
		} else {
			a = "1"
		}
		handler.SendResponse(c, a, days)
		return
	}

	if days == 14 {
		err = model.ChangeDeskmateStatus(DeskmateId)
		if err != nil {
			handler.SendResponse500(c, "未成功解除同桌关系", err) // 3.23
			return
		}
		handler.SendResponse(c, "恭喜你，成功完成本次同桌打卡!!!", nil)
		return
	}
}
