package services

import (
	"errors"
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/models"
	"golang.org/x/crypto/bcrypt"
)

func IsEmailRegistered(email string) (IsRegistered bool) {
	var auth models.Auth
	db.Where("email = ?", email).Find(&auth)
	if (auth == models.Auth{}) {
		IsRegistered = false
	} else {
		IsRegistered = true
	}
	return
}

func CreateUser(user models.Auth) {
	db.Create(&user)
}

func GetRoleWhileLogin(email string) (role string) {
	var user models.Auth
	db.Where("email = ?", email).Find(&user)
	role = user.Role
	return role
}

func IsEmailAndPasswordMatch(email string, passwd string) (isEmailAndPasswordMatch bool) {
	var user models.Auth
	db.Where("email = ?", email).Find(&user)
	result := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwd))
	if result == nil {
		isEmailAndPasswordMatch = true
	} else {
		isEmailAndPasswordMatch = false
	}
	return
}

func ResetUserPassword(email string, passwordHash string) (err error) {
	var user models.Auth
	db.Where("email = ?", email).Find(&user)
	if user.Email == "" {
		return errors.New("No such user in database")
	}
	user.Password = passwordHash
	db.Save(&user)
	return nil
}

func GetUserByEmail(email string) (user models.Auth, err error) {
	err = db.Where("email=?", email).Find(&user).Error
	return
}
