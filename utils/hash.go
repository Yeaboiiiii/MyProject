package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPassword(correct, entered string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(correct), []byte(entered))
	return err == nil
}
