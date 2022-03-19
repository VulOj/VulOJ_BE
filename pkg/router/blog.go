package router

import (
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/models"
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/services"
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetBlogsNumber(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"total_number": services.GetBlogsNumber(),
	})
}

func GetBlogs(context *gin.Context) {
	page, _ := strconv.Atoi(context.Query("page"))
	listSize, _ := strconv.Atoi(context.Query("blog_num"))
	blogs := services.GetBlogs(page, listSize, true)
	var results []map[string]interface{}
	for _, temp := range blogs {
		user, _ := services.GetUserByEmail(temp.AuthEmail)
		time, _ := util.ConvertShanghaiTimeZone(temp.PublishTimestamp)
		results = append(results, map[string]interface{}{
			"blog_id":           temp.ID,
			"title":             temp.Title,
			"content":           temp.Content,
			"auth_email":        temp.AuthEmail,
			"author_name":       user.Username,
			"publish_time":      time.String(),
			"organization_type": temp.OrganizationType,
			"organization_id":   temp.OrganizationID,
			"organization_name": temp.OrganizationName,
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": results,
	})
}

func GetBlog(context *gin.Context) {
	blogId, _ := strconv.Atoi(context.Param("blog_id"))
	blog, err := services.GetOneBlog(uint(blogId))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{`err`: err.Error()})
	} else {
		user, _ := services.GetUserByEmail(blog.AuthEmail)
		time, _ := util.ConvertShanghaiTimeZone(blog.PublishTimestamp)
		context.JSON(http.StatusOK, gin.H{
			"blog_id":           blog.ID,
			"title":             blog.Title,
			"content":           blog.Content,
			"auth_email":        blog.AuthEmail,
			"author_name":       user.Username,
			"publish_time":      time.String(),
			"organization_type": blog.OrganizationType,
			"organization_name": blog.OrganizationName,
			"organization_id":   blog.OrganizationID,
		})
		return
	}
}

func GetCommentOfBlog(context *gin.Context) {
	blogId, _ := strconv.Atoi(context.Param("blog_id"))
	comments, err := services.GetCommentOfBlog(uint(blogId))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{`err`: err.Error()})
	} else {
		var results []map[string]interface{}
		for _, temp := range comments {
			time, _ := util.ConvertShanghaiTimeZone(temp.PublishTimestamp)
			results = append(results, map[string]interface{}{
				"comment_content": temp.Content,
				"comment_time":    time.String(),
				"commenter_name":  temp.AuthName,
				"comment_id":      temp.ID,
			})
		}

		context.JSON(http.StatusOK, gin.H{
			"comments": results,
		})
	}
}

func AddComment(context *gin.Context) {
	blogId, _ := strconv.Atoi(context.PostForm("blog_id"))
	content := context.PostForm("comment_content")
	authEmail, _ := util.GetEmailFromToken(context)
	auth, _ := services.GetUserByEmail(authEmail)
	if content == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "请输入内容",
		})
		return
	}
	comment := models.Comment{
		BlogID:           uint(blogId),
		Content:          content,
		AuthEmail:        auth.Email,
		AuthName:         auth.Username,
		PublishTimestamp: time.Now(),
	}
	err := services.AddComment(&comment)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "发布评论失败",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"comment_id": comment.ID,
	})
	return
}

func AddBlog(context *gin.Context) {
	title := context.PostForm("title")
	content := context.PostForm("content")
	//organizationType := context.PostForm("organization_type")
	//organizationID, _ := strconv.Atoi(context.PostForm("organization_id"))
	//organizationName := context.PostForm("organization_name")
	authEmail, _ := util.GetEmailFromToken(context)
	if content == "" || title == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "请输入内容",
		})
		return
	}
	blog := models.BlogForbidden{
		Blog: models.Blog{
			AuthEmail:        authEmail,
			PublishTimestamp: time.Now(),
			Title:            title,
			Content:          content,
			//OrganizationType: organizationType,
			//OrganizationID:   uint(organizationID),
			//OrganizationName: organizationName,
		},
	}

	err, id := services.AddBlog(blog)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "发布文章失败",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"blog_id": id,
	})
	return
}

//只能删除自己发布的文章
func DeleteBlog(context *gin.Context) {
	blogId, _ := strconv.Atoi(context.PostForm("blog_id"))
	authEmail, _ := util.GetEmailFromToken(context)
	isDeleted, err := services.DeleteBlog(uint(blogId), authEmail)
	if isDeleted == false && err == nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "无权限删除文章",
		})
		return
	}
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "删除失败，请重试",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
	return
}

func GetBlogForbiddens(context *gin.Context) {
	page, _ := strconv.Atoi(context.Query("page"))
	listSize, _ := strconv.Atoi(context.Query("blog_num"))
	blogForbiddens := services.GetBlogForbiddens(page, listSize, true)
	var results []map[string]interface{}
	for _, temp := range blogForbiddens {
		user, _ := services.GetUserByEmail(temp.AuthEmail)
		time, _ := util.ConvertShanghaiTimeZone(temp.PublishTimestamp)
		results = append(results, map[string]interface{}{
			"blog_id":           temp.ID,
			"title":             temp.Title,
			"content":           temp.Content,
			"auth_email":        temp.AuthEmail,
			"author_name":       user.Username,
			"publish_time":      time.String(),
			"organization_type": temp.OrganizationType,
			"organization_id":   temp.OrganizationID,
			"organization_name": temp.OrganizationName,
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": results,
	})
}

func EnableBlog(context *gin.Context) {
	blogId, _ := strconv.Atoi(context.Param("blog_id"))
	blog, err := services.GetOneBlogForbidden(uint(blogId))
	if &blog == nil {
		context.Abort()
		context.JSON(http.StatusNotFound, gin.H{`msg`: "不存在此blog"})
	}
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{`err`: err.Error()})
	} else {
		err := services.EnableBlogAndDeleteForbiddenBlog(blog)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"msg": "数据库插入删除错误",
			})
			return
		} else {
			context.JSON(http.StatusOK, gin.H{
				"msg": "修改成功",
			})
		}
	}
}
