package card //修改名片头像，因为主要是以名片功能为主，这里就把头像直接放到名片功能里了
import (
	"Deskmate/model"
	"Deskmate/handler"
	"Deskmate/services"
	"Deskmate/services/connector"
	"Deskmate/services/user"
	"log"
	"os"
	"path"
	//"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 修改头像
// @Tags card
// @Description "修改名片头像"
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Body file form-Data file true "文件"
// @Success 200 {object} model.Card "{"mgs":"success"}"
// @Failure 400 "上传失败,请检查token与其他配置参数是否正确"
// @Router /card/avatar [post]
func ModifyUserProfile(c *gin.Context) {

	// temp := c.Request.Header.Get("id")
	// id, _ := strconv.Atoi(temp)

	/* 	temp, ok := c.Get("id")
	   	id := temp.(int)
	   	if !ok {
	   		c.JSON(401, gin.H{"message": "验证失败"})
	   	} */
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token) //这里用的是写在service里的解析token函数,并获取用户id即学号
	log.Println(id)
	PATH := "cards"
	if err != nil {
		// c.JSON(401, gin.H{"message": "Token Invalid"})
		handler.SendResponse401(c,"Token Invalid.",err) // 3.23
		return
	}

	//var cardInfo model.Card

	file, err := c.FormFile("file")

	if err != nil {
		/* c.JSON(400, gin.H{
			"msg": "上传失败1!",
		}) */
		handler.SendBadRequest(c, "上传失败", nil)
		return
	}

	filepath := "./"
	if _, err := os.Stat(filepath); err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(filepath, os.ModePerm)
		}
	}

	fileExt := path.Ext(filepath + file.Filename)

	//id1 := strconv.Itoa(id)

	file.Filename = id + "_" + services.GetRandomString(16) + fileExt

	filename := filepath + file.Filename

	if err := c.SaveUploadedFile(file, filename); err != nil {
		/* c.JSON(400, gin.H{
			"msg": "上传失败2!",
		}) */
		handler.SendBadRequest(c, "上传失败", err)
		return
	}

	// 删除原头像
	cardInfo, _ := model.GetCardInfo(id)
	if cardInfo.Path != "" && cardInfo.Sha != "" {
		connector.RepoCreate().Del(cardInfo.Path, cardInfo.Sha)
	}

	// 上传新头像
	Base64 := services.ImagesToBase64(filename)
	picUrl, picPath, picSha := connector.RepoCreate().Push(PATH,file.Filename, Base64)

	os.Remove(filename)

	var avatar model.Card
	avatar.UserId = id
	avatar.Avatar = picUrl
	avatar.Path = picPath
	avatar.Sha = picSha
	err0 := model.UpdateAvator(avatar)
	if picUrl == "" || err0 != nil {
		handler.SendBadRequest(c, "上传失败,请检查token与其他配置参数是否正确", err0)
		return
	}

	handler.SendResponse(c, "上传成功", map[string]interface{}{
		"url":  picUrl,
		"sha":  picSha,
		"path": picPath,
	})
}

// 七牛云图床
/* func Upload(c *gin.Context) {
	file,fileHeader,_:=c.Request.FormFile("file")

	fileSize := fileHeader.Size

	url,code:=model.UploadFile(file,fileSize)

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
		"url":url,
	})
}  */