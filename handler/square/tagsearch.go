package square

import (
	"Deskmate/model"

	"github.com/gin-gonic/gin"
)

// @Summary "搜索标签"
// @Description "在同桌广场中搜索标签返回对应名片"
// @Tags square
// @Accept json
// @Produce json
// @Param temp body model.Search true "temp"
// @Success 200 {object} model.BooksInfo "搜索成功"
// @Failure 401 "请重试"
// @Failure 404 "搜索不到"
// @Router /square/tag [get]
func TagSearch(c *gin.Context) {
	var a model.Search
	if err := c.BindJSON(&a); err != nil {
		c.JSON(401, gin.H{"message": "请重试"})
		return
	}
	result, err2 := model.GetCardByTag(a.Tag)
	//fmt.Println(temp.Content)
	if err2 != nil {
		c.JSON(404, gin.H{"message": "搜索不到"})
		return
	}
	c.JSON(200, result)
}
