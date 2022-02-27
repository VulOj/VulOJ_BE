package services

import "github.com/VulOJ/Vulnerable_Online_Judge_Project/models"

func CreateDirectory(Directory models.Directories) {
	db.Create(&Directory)
}
