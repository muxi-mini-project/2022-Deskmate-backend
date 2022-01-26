package card //1.24 第一次设置名片，后面都是修改名片

import (
	"Deskmate/handler"
	"Deskmate/model"
	"Deskmate/service/user"
	"log"

	"github.com/gin-gonic/gin"
)

// @Summary "设置名片"
// @Description "设置名片的昵称，标签，头像，用户宣言"
// @Tags card
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} []model.Organization "{"msg":"新建成功"}"
// @Failure 203 {object} error.Error "{"error_code":"20001","message":"Fail."}"
// @Failure 401 {object} error.Error "{"error_code":"10001","message":"Token Invalid."} 身份验证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"20001","message":"Fail."}or {"error_code":"00002","message":"Lack Param or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /api/v1/card [post]
func Create(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token) //这里用的是写在service里的解析token函数,并获取用户id即学号
	log.Println(id)

	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid"})
		return
	}
	var cardInfo model.Card
	cardInfo.UserId = id //这里直接从token里面调用学号放入名片，不需要在客户端手动输入
	cardInfo.Status = "无"
	if err := c.BindJSON(&cardInfo); err != nil { //BindJSON,把客户端输入的数据(如postman的json格式数据)存入cardInfo结构体中
		c.JSON(400, gin.H{
			"message": "Lack Param Or Param Not Satisfiable.", //缺少参数或参数不可满足。
		})
		return
	}
	result := model.DB.Create(&cardInfo)
	//result := model.ChangeCardInfo(cardInfo)
	if result != nil {
		c.JSON(400, "Fail.")
	}
	handler.SendResponse(c, "创建成功", nil)
}
