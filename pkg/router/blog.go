package router

import (
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
