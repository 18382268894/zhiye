/**
*FileName: token
*Create on 2018/11/7 下午4:04
*Create by mok
*/

package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"blog1/pkg/conf"
	"errors"
)


type MyClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewToken(username string)(tokenss string,err error){
	claim := &MyClaim{
		username,
		jwt.StandardClaims{
			Issuer:"blog1",
			ExpiresAt:time.Now().Unix()+conf.TokenMaxLife,
			IssuedAt:time.Now().Unix(),
			NotBefore:time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	tokenss,err = token.SignedString([]byte(conf.TokenSecret))
	return
}

func ParseToken(tokenss string)(claim *MyClaim,err error){
	token,err := jwt.ParseWithClaims(tokenss,&MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.TokenSecret),nil
	})
	if err != nil{
		return nil,err
	}
	claim,ok := token.Claims.(*MyClaim)
	if ok && token.Valid {
		return claim,nil
	}
	return nil,errors.New("cannot parse the token")
}
