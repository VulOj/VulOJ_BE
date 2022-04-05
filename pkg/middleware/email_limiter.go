package middleware

import (
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/consts"
	redis "github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/services"
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func EmailLimiter(apiName string) gin.HandlerFunc {
	return func(context *gin.Context) {
		email := context.PostForm("email")
		key := apiName + ":" + email
		isInBlackList := redis.IsInvolvedInBlackList(email)

		//To recognize which api in blacklist
		blacklist_key := consts.BLACK_LIST_PREFIX + apiName + ":" + email

		if isInBlackList {
			re, _ := redis.RedisClient.Get(blacklist_key).Int64()
			time_left := strconv.Itoa(int(30 - (util.GetTimeStamp()-re)/60))
			context.Abort()

			context.JSON(http.StatusTooManyRequests, gin.H{
				"msg": "请求过于频繁，请" + time_left + "分钟后再来:)",
			})
			return
		} else if re, _ := redis.RedisClient.Get(key).Int(); re >= 2 {
			context.Abort()
			redis.AddUserToBlackList(key, 8)
			context.JSON(http.StatusTooManyRequests, gin.H{
				"msg": "过于频繁请求!",
			})
			return
		} else {
			if re, _ := redis.RedisClient.Exists(key).Result(); re == 1 {
				redis.RedisClient.Incr(key)
			} else {
				redis.RedisClient.Set(key, 1, time.Minute)
			}
			context.Next()
		}

	}
}
