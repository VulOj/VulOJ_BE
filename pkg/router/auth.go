package router

import (
	"fmt"
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/models"
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/consts"
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/services"
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func HelloWorld(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Hello World!",
	})
	return
}

func Signup(c *gin.Context) {
	email := strings.ToLower(c.PostForm("email"))
	passwd := c.PostForm("password")
	verifyStr := c.PostForm("verify_code")
	name := c.PostForm("name")

	//rsa 2.0.0

	//check whether registered by email
	if services.IsEmailRegistered(email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "该用户已经注册",
		})
		return
	}

	//check whether VerifyCode Is Match To RegisterAccount
	if services.IsVerifyCodeMatchToRegisterAccount(verifyStr, email) {
		services.RemoveVerifyFromRedis(email)
		passwdHash := util.HashWithSalt(passwd)
		user := models.Auth{
			Email:             email,
			Password:          passwdHash,
			Username:          name,
			RegisterTimestamp: util.GetTimeStamp(),
			Role:              consts.USER,
		}
		services.CreateUser(user)
		token, _ := util.GenerateToken(email, consts.USER)
		c.SetCookie(consts.COOKIE_NAME, token, consts.EXPIRE_TIME_TOKEN, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"email": user.Email,
			"msg":   "注册成功",
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误，请重新输入验证码",
		})
		return
	}
}
func Login(context *gin.Context) {
	email := strings.ToLower(context.PostForm("email"))
	passwd := context.PostForm("password")

	role := services.GetRoleWhileLogin(email)
	token, _ := util.GenerateToken(email, role)

	if !services.IsEmailRegistered(email) {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "该账户未注册，请首先注册",
		})
		return
	}
	if services.IsEmailAndPasswordMatch(email, passwd) {
		context.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		return
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户名密码不匹配",
		})
		return
	}
}

func ChangePassword(context *gin.Context) {
	email := strings.ToLower(context.PostForm("email"))
	verifyCode := context.PostForm("verify_code")
	newPasswd := context.PostForm("new_password")

	if services.IsVerifyCodeMatchToRegisterAccount(verifyCode, email) {
		passwordHash := util.HashWithSalt(newPasswd)
		err := services.ResetUserPassword(email, passwordHash)
		services.RemoveVerifyFromRedis(email)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		} else {
			context.JSON(http.StatusOK, gin.H{
				"msg": "修改密码成功",
			})
			return
		}
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码过期或不正确",
		})
		return
	}
}

func GetMyselfInfo(context *gin.Context) {
	email, _ := util.GetEmailFromToken(context)
	user, _ := services.GetUserByEmail(email)
	context.JSON(http.StatusOK, gin.H{
		"name":  user.Username,
		"email": user.Email,
	})
	return
}

func Auth() gin.HandlerFunc {
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
			context.Next()
		} else {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{
				"msg": "无效token，请重新登录",
			})
			return
		}
	}
}

func SendVerifyCode(context *gin.Context) {
	email := strings.ToLower(context.PostForm("email"))
	isMailWell := strings.Compare(email, "2984252780@qq.com")
	if isMailWell == 0 {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "发送验证码失败，请重试，如果多次失败，请联系管理员",
		})
		fmt.Println("邮件服务被恶意调用")
		return
	}
	verifyCode := util.GenerateVerifyCode(consts.VERIFYCODE_LENGTH)
	services.StoreEmailAndVerifyCodeInRedis(verifyCode, email)

	if sendEmailSuccessful := services.SendEmail(email, verifyCode); sendEmailSuccessful {
		context.JSON(http.StatusOK, gin.H{
			"msg": "成功发送验证码，请注意查收",
		})
		return
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "发送验证码失败，请重试，如果多次失败，请联系管理员",
		})
		return
	}
}
