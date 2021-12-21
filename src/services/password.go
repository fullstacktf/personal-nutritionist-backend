package services

import "golang.org/x/crypto/bcrypt"

func GenerateHashPassword(password string) (string, error) {
	if string(password[0]) == "$" && len(password) > 30 {
		return password, nil
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckHashPassword(password, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
