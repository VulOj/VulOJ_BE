package services

import "github.com/VulOJ/Vulnerable_Online_Judge_Project/models"

func GetBlogsNumber() (number int) {
	var blog models.Blog
	db.Model(&blog).Count(&number)
	return number
}
