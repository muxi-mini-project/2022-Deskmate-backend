package user//登录用的

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

//const用来定义常量，常量是不可以被改变的
const (
	ErrorReasonServerBusy = "服务器繁忙"
	ErrorReasonReLogin    = "请重新登录"
)

type Jwt struct{
	StudentID string `json:"student_id"`
	jwt.StandardClaims
}

//user:
func VerifyToken(strToken string) (string,error) {
	//解析token
	token,err :=jwt.ParseWithClaims(strToken,&Jwt{},func(token *jwt.Token)(interface{},error) {
		return []byte("vinegar"),nil

	})
	if err!= nil{
		return "", errors.New(ErrorReasonServerBusy + ",或token解析失败")	
	}
	claims,ok :=token.Claims.(*Jwt)
	if !ok {
		return "",errors.New(ErrorReasonReLogin)
	}
	if err:=token.Claims.Valid();err!=nil{
		return "",errors.New(ErrorReasonReLogin)
	}
	return claims.StudentID,nil
}
