package services

import "github.com/VulOJ/Vulnerable_Online_Judge_Project/models"

func AddComment(comment *models.Comment) (err error) {
	err = db.Create(comment).Error
	return err
}
