package apply //创建同桌申请

import (
	"Deskmate/model"
	"Deskmate/services/user"
	"Deskmate/handler"

	"github.com/gin-gonic/gin"
)

// @Summary 同桌申请
// @Description 从名片页面像对方发出申请
// @Tags apply
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param object body model.Respondent true "申请对象的学号(id)"
// @Success 200  "{"msg":"success", "对象id":"string"}"
// @Failure 401  "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400  "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500  "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /apply [post]
func CreateApplication(c *gin.Context) {
	//swag里删除了500 @Failure 500  "{"error_code":"30001", "message":"database does not open successful"} 失败"
	token := c.Request.Header.Get("token")

	id, err := user.VerifyToken(token)

	if err != nil {
		// c.JSON(401, gin.H{"message": "验证失败"})
		handler.SendResponse401(c, "验证失败", err) // 3.21
		return
	}

	var application model.Applycation
	application.UserId1 = id

	var respondent model.Respondent

	if err := c.BindJSON(&respondent); err != nil { //从客户端读取申请对象的学号
		/* c.JSON(400, gin.H{
			"message": "Lack Param Or Param Not Satisfiable.", //缺少参数或参数不可满足。
		}) */
		handler.SendBadRequest(c,"Lack Param Or Param Not Satisfiable.",err)
		return
	}

	if respondent.StudentID == "" {
		handler.SendBadRequest(c, "Lack Param Or Param Not Satisfiable.", nil) //4.8
		return
	}

	result, err := model.GetUserStatus(respondent.StudentID) //判断此时申请对象是否有同桌

	if err != nil {
		/* c.JSON(401, gin.H{
			"message": "未查询到申请对象是否有同桌.",
		}) */
		handler.SendResponse500(c,"未查询到申请对象是否有同桌",err)
		return //用return直接返回，不需要用If语句
	}

	if result == "1" {
		// c.JSON(200, "对方已有同桌，发送申请失败")
		handler.SendResponse(c,"对方已有同桌，发送申请失败",nil)
		return
	}

	//4.7 修复了能像自己申请同桌的bug，经过测试发现result是能够正常查询的
	if id == respondent.StudentID { 
		handler.SendResponse(c,"你不能成为自己的同桌哦！",nil)
		return
	}

	if result == "0" {
		var apply model.Applycation
		apply.UserId1 = id
		apply.UserId2 = respondent.StudentID
		if apply.UserId2 == " "{
			handler.SendBadRequest(c,"申请对象不存在！",nil)
			return
		}
		apply.Result = " "
		if err := model.DB.Create(&apply).Error; err != nil {
			/* c.JSON(200, gin.H{
				"msg":   "create fail",
				"error": err,
			}) */
			handler.SendResponse500(c,"发送申请失败！",err)
			return
		}

		/* c.JSON(200, gin.H{
			"msg":           "success",
			"respondent_id": respondent.StudentID,
		}) */
		handler.SendResponse(c, "发送申请成功！", respondent.StudentID)
		return
	}
}

/*func CreateApplication(user1 string, user2 string) (string, error) {//user1为发出申请者，user2为收到申请者
	var apply model.Apply
	apply.UserId1 = user1
	apply.UserId2 = user2
	if err := DB.Create(&apply).Error;err !=nil{
		return "创建申请失败",err
	}
	return "创建申请成功",nil
}*/
