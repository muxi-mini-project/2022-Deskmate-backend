package dailyrecord //提前中断同桌关系

import (
	"Deskmate/model"
	"Deskmate/handler"
	"Deskmate/services/user"

	"github.com/gin-gonic/gin"
)

// @Summary 强制中断同桌
// @Description "强制中断当前同桌"
// @Tags dailyrecord
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 "强制中断成功"
// @Failure 401 "身份验证失败 重新登录"
// @Failure 400 "强制中断关系失败"
// @Failure 500 "用户更新状态失败"
// @Router /dailyrecord/end [put]
func BreakRelation(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		// c.JSON(401, gin.H{"message": "验证失败"})
		handler.SendResponse401(c,"Token Invalid.",nil) // 3.23
		return
	}
	//开始前先判断此时是否有同桌，如果没有，则告诉用户前往同桌广场寻找同桌
	result, err := model.GetUserStatus(id) //判断此时申请对象是否有同桌
	if err != nil {
		/* c.JSON(401, gin.H{
			"message": "未查询到申请对象是否有同桌.",
		}) */
		handler.SendResponse500(c,"未查询到申请对象是否有同桌.",err) // 3.23
		return //用return直接返回，不需要用If语句
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

	/* sid1, sid2, err := model.GetPartnerId(DeskmateId)
	if err != nil {
		handler.SendResponse401(c,"未查询到相关成员",err) 
		return
	} */
	sid1 := id //本人的id
	sid2 , err :=model.GetPartner(DeskmateId,id) //同桌的id
	if err != nil {
		handler.SendResponse500(c,"未查询到相关成员！",err) // 3.23
		return
	} 

	err = model.ChangeStatusAgain(sid1)

	if err != nil {
		// c.JSON(500, gin.H{"message": "用户更新状态失败"})
		handler.SendResponse500(c,"用户更新状态失败",err) // 3.23
		return
	}

	err = model.ChangeStatusAgain(sid2)

	if err != nil {
		// c.JSON(500, gin.H{"message": "用户更新状态失败"})
		handler.SendResponse500(c,"用户更新状态失败",err) // 3.23
		return
	}

	err = model.ChangeDeskmateStatus(DeskmateId)
	if err != nil {
		handler.SendResponse500(c,"解除同桌关系失败",err) // 3.23
		return
	}

	// c.JSON(200, "成功解除同桌关系!")
	handler.SendResponse(c, "成功解除同桌关系!", nil) // 3.23
}
