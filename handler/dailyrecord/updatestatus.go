package dailyrecord //用于更新每天的同桌状态，如果不符合则立即结束，如果符合则天数加一，如果14天已满则可以再次发送申请是否继续成为同桌

import (
	"Deskmate/model"
	"Deskmate/services/user"

	"github.com/gin-gonic/gin"
)

// @Summary 更新打卡
// @Description "每日更新打卡天数"
// @Tags dailyrecord
// @Accept json
// @Produce json
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
		c.JSON(401, gin.H{"message": "验证失败"})
	}
	DeskmateId, err := model.GetDeskmateId(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "未查询到对应同桌关系id",
		})
		return
	}

	sid1, sid2, err := model.GetPartnerId(DeskmateId)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "未查询到相关成员",
			"error":   err,
		})
		return
	}

	days, err := model.GetDays(DeskmateId)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "未查询到打卡天数",
			"error":   err,
		})
		return
	}

	if days == 14 {
		c.JSON(200, gin.H{
			"message": "你已成功完成本次同桌打卡!",
		})
		return
	}
	if days == 0 {
		c.JSON(200, gin.H{
			"message": "你还为开始打卡哦!",
		})
		return
	}

	//判断两人是否都在昨天打卡
	judgement1, err := model.GudgeYesterday(sid1, DeskmateId)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "未查询用户昨天的记录",
			"error":   err,
		})
		return
	}
	judgement2, err := model.GudgeYesterday(sid2, DeskmateId)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "未查询用户昨天的记录",
			"error":   err,
		})
		return
	}
	if judgement1 == "false" || judgement2 == "false" {
		err = model.ChangeDeskmateStatus(DeskmateId)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "未成功解除同桌关系",
				"error":   err,
			})
			return
		}
		c.JSON(400, gin.H{
			"message": "很遗憾，有人未按照要求及时打卡，本次同桌打卡终止!",
		})
	}

	if judgement1 == "false" && judgement2 == "false" {
		days++
		err = model.UpdateDays(DeskmateId, days)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "更新天数错误",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "更新成功",
		})
	}
}
