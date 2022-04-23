package square

import (
	"Deskmate/model"
	"Deskmate/handler"
	"github.com/gin-gonic/gin"
)

// @Summary 搜索标签
// @Description "在同桌广场中搜索标签返回对应名片"
// @Tags square
// @Accept application/json
// @Produce application/json
// @Param object body model.Tag true "输入要搜索的标签"
// @Success 200 {object} []model.Card "搜索成功"
// @Failure 401 "Lack Param Or Param Not Satisfiable."
// @Failure 404 "搜索失败"
// @Router /square/tag [post]
func TagSearch(c *gin.Context) {
	var a model.Search
	if err := c.BindJSON(&a); err != nil {
		// c.JSON(401, gin.H{"message": "请重试"})
		handler.SendResponse401(c, "Lack Param Or Param Not Satisfiable.",err)
		return
	}
	result, err2 := model.GetCardByTag(a.Tag)
	//fmt.Println(temp.Content)
	if err2 != nil {
		// c.JSON(404, gin.H{"message": "搜索不到"})
		handler.SendResponse404(c, "搜素失败", err2)
		return
	}


	// c.JSON(200, result)
	if result == nil {
		handler.SendResponse(c,"搜不到此tag",result)
	}
	handler.SendResponse(c, "搜索成功", result)
}
