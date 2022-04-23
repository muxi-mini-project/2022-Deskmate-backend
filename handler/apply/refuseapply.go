package apply // 3.22 增加拒绝申请
import (
	"Deskmate/model"
	"Deskmate/services/user"
	"Deskmate/handler"
	"github.com/gin-gonic/gin"
)

// @Summary 拒绝同桌申请
// @Description 用户拒绝接受的同桌申请
// @Tags apply
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param object body model.Respondent true "要拒绝申请的同学学号"
// @Success 200 "{"msg":"success"}"
// @Failure 401 "{"msg":"confirm faided"}"
// @Router /apply/refuse [put]
func ApplicationRefuse(c *gin.Context) {
	//swag删除了400和500 @Failure 400 "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
	// @Failure 500 "{"error_code":"30001", "message":"Fail."} 失败
	token := c.Request.Header.Get("token")

	id, err := user.VerifyToken(token)

	if err != nil {
		// c.JSON(401, gin.H{"message": "验证失败"})
		handler.SendResponse401(c, "身份验证失败!", err) // 3.21
		return
	}

	// rid := c.Query("respondent_id")

	
	var respondent model.Respondent

	if err := c.BindJSON(&respondent); err != nil { //从客户端读取需要确认的申请对象的学号
		handler.SendBadRequest(c,"Lack Param Or Param Not Satisfiable.",err)
		return
	}

	rid := respondent.StudentID // 3.22 这里算是一个比较大的bug，原来的query语句好像根本无法读入数据
	
	err = model.RefuseApplication(id, rid)

	if err != nil {
		// c.JSON(401, gin.H{"message": "申请结果确认失败"})
		handler.SendResponse401(c, "数据库操作失败！", err) // 3.21
		return
	}

	err = model.ChangeStatus(id)

	if err != nil {
		// c.JSON(401, gin.H{"message": "用户更新状态失败"})
		handler.SendResponse401(c, "用户更新状态失败！", err) // 3.21
		return
	}

	err = model.ChangeStatus(rid)

	if err != nil {
		// c.JSON(401, gin.H{"message": "用户更新状态失败"})
		handler.SendResponse401(c, "用户更新状态失败！", err) // 3.21
		return
	}

	handler.SendResponse(c, "拒绝该申请成功!", "") // 3.22
}