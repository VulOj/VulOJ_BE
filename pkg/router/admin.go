package router

import (
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/consts"
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/services"
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRoot() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("token")
		claim, err := util.GetClaimFromToken(tokenString)
		if err != nil {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{
				"msg": "未认证，请先登录",
			})
			return
		}
		email := claim.(jwt.MapClaims)["email"].(string)

		//using assert is very dangerous
		tokenTimeStamp := claim.(jwt.MapClaims)["timeStamp"].(float64)
		time := util.GetTimeStamp() - int64(tokenTimeStamp)
		if time > consts.EXPIRE_TIME_TOKEN {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Token过期，请重新登录",
			})
		}

		if services.IsEmailRegistered(email) {
			role := claim.(jwt.MapClaims)["role"].(string)
			if role != "root" {
				context.Abort()
				context.JSON(http.StatusUnauthorized, gin.H{
					"msg": "非root用户",
				})
			} else {
				context.Next()
			}
		} else {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{
				"msg": "无效token，请重新登录",
			})
			return
		}
	}
}
