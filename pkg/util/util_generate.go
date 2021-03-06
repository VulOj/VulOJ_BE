package util

import (
	"crypto/rand"
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/consts"
	"github.com/dgrijalva/jwt-go"
	"io"
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

//remind that params which store tokenstr must name token:)
func GenerateVerifyCode(length int) (verifyString string) {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	verifyString = string(b)
	return

}
