package user

import(
	"Deskmate/handler"
	"Deskmate/model"
	"Deskmate/service/user"
	"encoding/base64"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"//json-web-token
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context){
	var p model.User
	if err := c.BindJson(&p); err != nil{
		c.JSON(400,gin.H{"message":"Lack Param Or Param Not Satisfiable."})
		return
	}
	if p.StudentID == "" {
		c.JSON(400,gin.H{"message":"Lack Param or Param Not Satisfiable."})
		return
	}
	pwd :=p.Password
}