package apply //游览收到的同桌申请

import (
	"Deskmate/handler"
	"Deskmate/model"
	"Deskmate/services/user"

	"github.com/gin-gonic/gin"
)

// @Summary 显示收到的同桌申请
// @Description 点击 同桌申请 时调用
// @Tags apply
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200  "{"msg":"success"}"
// @Failure 401  "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400  "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500  "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /apply [get]
func ViewApplication(c *gin.Context) {
	token := c.Request.Header.Get("token")

	id, err := user.VerifyToken(token)

	if err != nil {
		// c.JSON(401, gin.H{"message": "验证失败"})
		handler.SendResponse401(c, "身份验证失败!", err) // 3.21
		return
	}
	Applyinformation, err := model.GetMyApply(id)
	if err != nil {
		// c.JSON(404, gin.H{"message": "获取失败"})
		handler.SendBadRequest(c, "获取失败", err) // 3.21
		return
	}

	var Cardsinfors []interface{}
	var Cardsinfor model.Card
	for _,m := range Applyinformation{//i是循环次数,这里用不到就省略了，m是Applyinformation[i];
		Cardsinfor, err = model.GetCardInfo(m.UserId1)
		if err != nil {
			handler.SendResponse500(c,"获取名片信息失败！",err)
		}
		Cardsinfors = append(Cardsinfors, Cardsinfor)
	}
	/* applicantId, err := model.GetMyApplyUserId1(id)
	if err != nil {
		handler.SendBadRequest(c, "获取申请人id失败", err) //4.20
		return
	}
	len := len(applicantId)
	var Cardsinfor []model.Card
	for i := 0; i < len; i++ {
		Cardsinfor[i], err = model.GetCardInfo(applicantId[i])
	} */
	// c.JSON(200, Applyinformation)
	//handler.SendResponse(c, "获取成功!", Applyinformation)
	handler.SendResponse(c, "获取成功!", Cardsinfors)
}
