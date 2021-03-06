package dailyrecord //同桌打卡，查看同桌名片

import (
	"Deskmate/model"
	"Deskmate/handler"
	"Deskmate/services/user"

	"github.com/gin-gonic/gin"
)

// @Summary 查看同桌名片
// @Description  点击查看同桌名片
// @Tags dailyrecord
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 "成功"
// @Failure 401  "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400  "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500  "{"error_code":"30001", "message":"Fail."} 服务器错误"
// @Router /dailyrecord/card [get]
func Cardinfo(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		// c.JSON(401, gin.H{"message": "验证失败"})
		handler.SendResponse401(c,"Token Invalid.",nil) // 3.23
		return
	}

	// 开始前先判断此时是否有同桌，如果没有，则告诉用户前往同桌广场寻找同桌
	result, err := model.GetUserStatus(id) //判断此时申请对象是否有同桌
	if err != nil {
		/* c.JSON(401, gin.H{
			"message": "未查询到申请对象是否有同桌.",
		}) */
		handler.SendResponse500(c,"未查询到申请对象是否有同桌.",err) // 3.23
		return // 用return直接返回，不需要用If语句
	}

	if result == "0" {
		// c.JSON(200, "你还没有同桌，请前往同桌广场寻找同桌!")
		handler.SendResponse(c,"你还没有同桌，请前往同桌广场寻找同桌！",nil) // 3.23
		return
	}

	DeskmateId, err := model.GetDeskmateId(id)
	if err != nil {
		/* c.JSON(401, gin.H{
			"message": "未查询到对应同桌关系id",
		}) */
		handler.SendResponse401(c,"未查询到对应同桌关系id",err) // 3.23
		return
	}

	/* sid1, sid2, err := model.GetPartnerId(DeskmateId) //这里函数加了个id
	if err != nil {
		c.JSON(401, gin.H{
			"message": "未查询到相关成员",
			"error":   err,
		}) 
		handler.SendResponse500(c,"未查询到相关成员！",err) // 3.23
		return
	} */

	//查询到的两个id一个是自己的一个是自己的同桌的，哪个和自己的Id一样则另一个就是同桌的
	/* if sid1 == id {
		id = sid2
	} else {
		id = sid1
	} */

	id , err =model.GetPartner(DeskmateId,id)
	if err != nil {
		handler.SendResponse500(c,"未查询到相关成员！",err) // 3.23
		return
	} 

	Userinformation, err := model.GetCardInfo(id)

	if err != nil {
		// c.JSON(404, gin.H{"message": "获取失败"})
		handler.SendResponse500(c,"获取失败",err) // 3.23
		return
	}
	// c.JSON(200, Userinformation)
	handler.SendResponse(c, "success", Userinformation) // 3.23
}
