package services

import (
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/consts"
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/util"
	"time"
)

func IsInvolvedInBlackList(email_aes string) (isInvolved bool) {
	re, _ := RedisClient.Exists(consts.BLACK_LIST_PREFIX + email_aes).Result()
	if re == 1 {
		isInvolved = true
	} else {
		isInvolved = false
	}
	return
}

func AddUserToBlackList(email_aes string, mute_time float64) {
	RedisClient.Set(consts.BLACK_LIST_PREFIX+email_aes, util.GetTimeStamp(), time.Duration(mute_time)*time.Hour)
}
func RemoveFromBlackList(email_aes string) {
	RedisClient.Del(consts.BLACK_LIST_PREFIX + email_aes)
}
