package util

import "golang.org/x/crypto/bcrypt"

// A Hash function using salt with bcrypt libriry to hash password
func HashWithSalt(plainText string) (HashText string) {

	hash, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.MinCost)
	CheckError(err)
	HashText = string(hash)
	return
}
