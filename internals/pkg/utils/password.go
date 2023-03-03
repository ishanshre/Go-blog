package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	encPass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(encPass), nil
}

func VerifyPassword(hashedPassword, reqestPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(reqestPassword))
}
