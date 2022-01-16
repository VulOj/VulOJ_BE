package services

import (
	"fmt"
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/consts"
)

func IsVerifyCodeMatchToRegisterAccount(verifyCode string, email string) (IsMatch bool) {
	re := RedisClient.Get(consts.REDIS_VERIFY_CODE_SUFFIX + email)
	fmt.Println("\n\n", re.Val())
	if re.Val() == verifyCode && re.Val() != "" {
		IsMatch = true
	} else {
		IsMatch = false
	}
	return
}
func RemoveVerifyFromRedis(email string) {
	RedisClient.Del(consts.REDIS_VERIFY_CODE_SUFFIX + email)
}
