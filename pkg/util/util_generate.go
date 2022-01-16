package util

import (
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/consts"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(email string, role string) (tokenString string, err error) {
	claims := jwt.MapClaims{
		"email":     email,
		"timeStamp": GetTimeStamp(),
		"role":      role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(consts.TOKEN_SCRECT_KEY))
	return
}
