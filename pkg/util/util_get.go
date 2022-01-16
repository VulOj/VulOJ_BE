package util

import (
	"errors"
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/consts"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func GetTimeStamp() (t int64) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t = time.Now().In(loc).Unix()
	return
}

func GetEmailFromToken(c *gin.Context) (email string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("您未登录，请登陆后查看")
		}
	}()
	tokenStr := c.GetHeader("token")
	claim, _ := GetClaimFromToken(tokenStr)
	email = claim.(jwt.MapClaims)["email"].(string)
	return
}

func GetClaimFromToken(tokenString string) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return []byte(consts.TOKEN_SCRECT_KEY), err
	})
	if err != nil {
		return nil, err
	} else {
		claims = token.Claims.(jwt.MapClaims)
		return claims, nil
	}
}
