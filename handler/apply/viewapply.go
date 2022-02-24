package apply //游览收到的同桌申请

import (
	"Deskmate/model"
	"Deskmate/services/user"

	"github.com/gin-gonic/gin"
)

// @Summary 显示收到的同桌申请
// @Description 点击 同桌申请 时调用
// @Tags apply
// @Accept json
// @Produce json
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
		c.JSON(401, gin.H{"message": "验证失败"})
	}
	Applyinformation, err := model.GetMyApply(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
	}
	c.JSON(200, Applyinformation)
}
