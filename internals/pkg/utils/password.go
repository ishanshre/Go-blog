package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// hash the plain password and returns hashed password
	encPass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(encPass), nil
}

func VerifyPassword(hashedPassword, reqestPassword string) error {
	// compare and verify hash password with plain password
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(reqestPassword))
}
