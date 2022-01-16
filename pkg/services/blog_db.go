package services

import "github.com/VulOJ/Vulnerable_Online_Judge_Project/models"

func GetBlogsNumber() (number int) {
	var blog models.Blog
	db.Model(&blog).Count(&number)
	return number
}

func GetBlogs(page int, listSize int, isDescend bool) (blogs []models.Blog) {
	if isDescend {
		db.Offset(page * listSize).Limit(listSize).Order(" publish_timestamp desc").Find(&blogs)
		/*db.Raw("SELECT blog.* ...起别名, auth_blog.username as ...
		FROM blog
		left join  auth_blog
		on blog.auth_email = auth_user.email
		order by  publish_timestamp desc
		limit ? , ?", page*listSize,listSize).Find(&result)
		result可以是map[string]interface{}
		*/
		//需要覆盖time
	} else {
		db.Offset(page * listSize).Limit(listSize).Find(&blogs)
	}
	return blogs
}

func GetOneBlog(blogId uint) (blog models.Blog, err error) {
	err = db.Where("id = ?", blogId).Find(&blog).Error
	return
}

func GetCommentOfBlog(blogId uint) (comments []models.Comment, err error) {
	err = db.Where("blog_id = ?", blogId).Find(&comments).Error
	return
}

//默认先加入blogForbiddens表不显示，审核之后插入到blog表显示
func AddBlog(blog models.BlogForbidden) (err error, id uint) {
	err = db.Create(&blog).Error
	id = blog.ID
	return
}

func DeleteBlog(blogId uint, authEmail string) (isDeleted bool, err error) {
	var blog models.Blog
	db.Where("id = ?", blogId).First(&blog)
	if blog.AuthEmail != authEmail {
		isDeleted = false
		return
	}
	err = db.Where("id = ?", blogId).Delete(&blog).Error
	if err != nil {
		isDeleted = false
		return
	}
	return true, nil
}
