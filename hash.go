package userservice

import "golang.org/x/crypto/bcrypt"

type BcryptPasswordManager struct{}

func (BcryptPasswordManager) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (BcryptPasswordManager) CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
