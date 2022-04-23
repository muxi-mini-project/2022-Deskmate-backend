package apply //确认收到的同桌申请

import (
	"Deskmate/model"
	"Deskmate/services/user"
	"Deskmate/handler"
	"github.com/gin-gonic/gin"
)

// @Summary 确认同桌申请,进入同桌打卡
// @Description 用户确认接受的同桌申请
// @Tags apply
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param object body model.Respondent true "要同意申请的同学学号"
// @Success 200 "{"msg":"success"}"
// @Failure 401 "{"msg":"confirm faided"}"
// @Router /apply [put]
// 如果两位同学有多次申请，则在每一次完成确认后将前一次申请删除(未实现),这里可以在游览申请时直接调出结果未知的申请，"如果成功，则将对应申请的状态(result)改为同意(1)"
func ApplicationConfirm(c *gin.Context) {
	//swag删除了400和500 @Failure 400 "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
	// @Failure 500 "{"error_code":"30001", "message":"Fail."} 失败
	token := c.Request.Header.Get("token")

	id, err := user.VerifyToken(token)

	if err != nil {
		// c.JSON(401, gin.H{"message": "验证失败"})
		handler.SendResponse401(c, "身份验证失败!", nil) // 3.21
		return
	}

	// rid := c.Query("respondent_id")

	var respondent model.Respondent

	if err := c.BindJSON(&respondent); err != nil { //从客户端读取需要确认的申请对象的学号
		handler.SendBadRequest(c,"Lack Param Or Param Not Satisfiable.",nil)
		return
	}

	rid := respondent.StudentID // 3.22 这里算是一个比较大的bug，原来的query语句好像根本无法读入数据
	
	err = model.ConfirmApplication(id, rid)

	if err != nil {
		// c.JSON(401, gin.H{"message": "申请结果确认失败"})
		handler.SendResponse401(c, "数据库操作失败！", nil) // 3.21
		return
	}

	err = model.ChangeStatus(id)

	if err != nil {
		// c.JSON(401, gin.H{"message": "用户更新状态失败"})
		handler.SendResponse401(c, "你的状态更新失败！", nil) // 3.21
		return
	}

	err = model.ChangeStatus(rid)

	if err != nil {
		// c.JSON(401, gin.H{"message": "用户更新状态失败"})
		handler.SendResponse401(c, "对方状态更新失败！", nil) // 3.21
		return
	}

	//这里是在确认后两人成为同桌记录到数据库
	var dailyrecord model.Dailyrecords
	dailyrecord.UserId1 = id
	dailyrecord.UserId2 = rid
	dailyrecord.Status = "进行中"
	if err := model.DB.Create(&dailyrecord).Error; err != nil {
		/* c.JSON(200, gin.H{
			"msg":   "create dailyrecord fail",
			"error": err,
		}) */
		handler.SendBadRequest(c, "create dailyrecord fail", err) // 3.21
		return
	}

	/* c.JSON(200, gin.H{
		"msg": "confirm success",
	}) */
	handler.SendResponse(c, "确认同桌关系成功!", "") // 3.21
}
