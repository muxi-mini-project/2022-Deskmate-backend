package apply //创建同桌申请

import (
	"Deskmate/model"
	"Deskmate/service/user"

	"github.com/gin-gonic/gin"
)

// @Summary 同桌申请
// @Description 从名片页面像对方发出申请
// @Tags apply
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param tmprequirement body model.VeRequirePassenger true "乘客需要填写的信息,注意年月日需要以xx年xx月xx日的形式填写,status表示该订单是否完成,1为未完成,2为已完成"
// @Success 200 {object} model.Res "{"msg":"success", "pid":"string"}"
// @Failure 401 {object} handler.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} handler.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} handler.Error "{"error_code":"30001", "message":"database does not open successful"} 失败"
// @Router /app;y [post]
func CreateApplication(c *gin.Context) {
	token := c.Request.Header.Get("token")

	id, err := user.VerifyToken(token)

	if err != nil {
		c.JSON(401, gin.H{"message": "验证失败"})
	}

	var application model.Apply
	application.UserId1 = id
	var respondent model.Respondent

	if err := c.BindJSON(&respondent); err != nil { //从客户端读取申请对象的学号
		c.JSON(400, gin.H{
			"message": "Lack Param Or Param Not Satisfiable.", //缺少参数或参数不可满足。
		})
		return
	}

	result, err := model.GetUserStatus(respondent.StudentID) //判断此时申请对象是否有同桌

	if err != nil {
		c.JSON(401, gin.H{
			"message": "未查询到申请对象是否有同桌.",
		})
	} else if result == "有" {
		c.JSON(200, "对方已有同桌，发送申请失败")
	} else {
		var apply model.Apply
		apply.UserId1 = id
		apply.UserId2 = respondent.StudentID
		if err := model.DB.Create(&apply).Error; err != nil {
			c.JSON(200, gin.H{
				"msg":   "create fail",
				"error": err,
			})
		}
		c.JSON(200, gin.H{
			"msg":  "success",
			"对象id": respondent.StudentID,
		})
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
