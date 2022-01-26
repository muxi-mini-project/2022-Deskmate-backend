package square //同桌广场的名片信息流，即从数据库中查询名片

import (
	"Deskmate/model"

	"github.com/gin-gonic/gin"
)

// @Summary "同桌广场"
// @Description "显示名片数据流"
// @Tags square
// @Accept json
// @Produce json
// @Success 200 {object} []model.Card "获取成功"
// @Router /square [get]
func ViewSquare(c *gin.Context) {
	Cardsinfor, err := model.GetCardsInfo()
	if err != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
	}
	c.JSON(200, Cardsinfor)
}
