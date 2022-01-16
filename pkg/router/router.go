package router

import (
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "access-control-allow-origin", "token"},
		AllowCredentials: false,
	}
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	api := r.Group("/api")
	{
		user := api.Group("/auth")
		user.Use(cors.Default())
		{
			user.GET("/hello", HelloWorld)
			user.POST("/signup", Signup)
			user.POST("/login", Login)
			user.POST("/changePassword", ChangePassword)
			user.GET("/myself", GetMyselfInfo)
			//	user.POST("/sendVerifyCode",
			//		middleware.IpLimiter("post/sendVerifyApi",2),
			//		middleware.EmailLimiter("email"),
			//		SendVerifyCode)
			//	user.POST("/changePasswordEmail", ChangePasswordByEmail)
			//	user.POST("/changePasswordVerifyCode", ChangePasswordVerifyCode)
			user.Use(Auth())
			//	user.GET("/getEnterpriseOfUser",GetEnterprisesOfUser)
			//	user.GET("/getProjectsOfUser",GetProjectsOfUser)
			//}
			//
			blog := api.Group("/blog")
			blog.Use(cors.Default())
			{
				blog.GET("/getBlogNumber", GetBlogsNumber)
				blog.GET("", GetBlogs)
				blog.GET("/detail/:blog_id", GetBlog)
				blog.GET("/comment/:blog_id", GetCommentOfBlog)
				blog.Use(Auth())
				{
					blog.POST("/comment", AddComment)
					blog.POST("", middleware.IpLimiter("post/add_blog", 5), AddBlog)
					blog.DELETE("", DeleteBlog)
				}
			}
			//
			//admin:=api.Group("/admin")
			//admin.Use(cors.Default(),AuthRoot())
			//{
			//	blog.POST("/enableBlog/:blog_id",EnableBlog)
			//	blog.GET("/getBlogForbidden",GetBlogForbiddens)
			//}
			//

		}
	}

}
