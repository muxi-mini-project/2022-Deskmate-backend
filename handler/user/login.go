package user

import (
	"Deskmate/handler"
	"Deskmate/model"
	"Deskmate/service/user"
	"encoding/base64"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go" //json-web-token，用来写token的
	"github.com/gin-gonic/gin"
)

//@Summary 用户登录
//@Tags user
//@Description 一站式登录
//@Accep applica/json
//@Produc applic/json
//@Param object body model.User trun "登录用户信息"
//Success 200{object} Token "将student_id作为token保留"
// @Success 200{object} handler.Response "{":"将student_id作为token保留"}"
// @Failure 401{object} error.Error "{"error_code":"10001","message":"password or account wrong."} 身份认证失败 重新登录"
// @Failure 400{object} error.Error "{"error_code":"20001","message":"Fail."} or {"error_code":"00002","message":"Lack Param Or Param Not Satisfiable."}"
// @Faliure 500{object} error.Error "{"error_code":"30001","message":"Fail."} 失败"
// @Router /user/login [post]
func Login(c *gin.Context) {
	var p model.User //model.User就是我们建的users模型表
	//var u model.Card //model.Card就是我们建的Cards模型表
	if err := c.BindJSON(&p); err != nil {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	if p.StudentID == "" {
		c.JSON(400, gin.H{"message": "Lack Param or Param Not Satisfiable."})
		return
	}
	pwd := p.PassWord
	//首次登录 验证一站式
	//判断是否首次登录，如果是第一次，则将信息录入数据库，如果从数据库中能查询到改信息，说明不是第一次登录
	result := model.DB.Where("student_id = ?", p.StudentID).First(&p)
	if result.Error != nil {
		//_, err := model.GetUserInfoFormOne(p.StudentID, pwd) //这个是model/studentinfo.go中的函数，是学长写好的一站式登录，返回一个用户信息的结构体
		userinfo, err := model.GetUserInfoFormOne(p.StudentID, pwd)
		if err != nil {
			c.JSON(401, "Password or account wrong.")
			return
		}
		//对用户信息初始化，通过学长的包可以直接爬取相关学生信息，比如姓名(Name)，学院(DeptName)，学号(Username)，本科生(UsertypeName)
		//p.Name = "请修改昵称"
		p.Name = userinfo.User.Name        //获取学生姓名
		p.College = userinfo.User.DeptName //获取学院名称
		s := userinfo.User.Username        //学号
		s = s[0:4]
		p.Grade = s //年级对应学号的前四位
		//p.Major = userinfo.User.UsertypeName//这里想读取具体的专业没有找到对应的对象
		//对密码进行base64加密
		p.PassWord = base64.StdEncoding.EncodeToString([]byte(p.PassWord))
		model.DB.Create(&p)
	} else {
		//在数据库中解密比较
		password, _ := base64.StdEncoding.DecodeString(p.PassWord)

		if string(password) != pwd {
			c.JSON(401, "Password or account is wrong.")
			return
		}
	}
	claims := &user.Jwt{StudentID: p.StudentID}

	claims.ExpiresAt = time.Now().Add(200 * time.Hour).Unix()
	claims.IssuedAt = time.Now().Unix()

	var Secret = "vinegar" //加醋

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		log.Println(err)
	}

	handler.SendResponse(c, "将student_id作为token保留", signedToken)
}