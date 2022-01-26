package token //好像原来就写了token的，这个是hjj写的，先放这里 1.24

import (
	"errors"
	"fmt"
	"log"

	//"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//phone唯一对应用户了，不需要获取用户id
//生成token与验证

type jwtClaims struct {
	jwt.StandardClaims        //jwt-go包预定义的一些字段
	Phone              string `json:"phone"`
}

var (
	key        = "miniProject"
	ExpireTime = 604800 //token过期时间
)

//我自己往token里写进去的只有phone
func GenerateToken(phone string) string {
	claims := &jwtClaims{
		Phone: phone,
	}
	//签发者和过期时间
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	singedToken, err := genToken(*claims)
	if err != nil {
		log.Print("produceToken err:")
		fmt.Println(err)
		return ""
	}
	return singedToken
}

func genToken(claims jwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	singedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return singedToken, nil
}

//验证token
func VerifyToken(token string) (string, error) {
	TempToken, err := jwt.ParseWithClaims(token, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return "", errors.New("token解析失败")
	}
	claims, ok := TempToken.Claims.(*jwtClaims)
	if !ok {
		return "", errors.New("发生错误")
	}
	if err := TempToken.Claims.Valid(); err != nil {
		return "", errors.New("发生错误")
	}
	fmt.Println(claims.Phone)
	return claims.Phone, nil
}
