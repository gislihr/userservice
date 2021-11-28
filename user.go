package userservice

import (
	"fmt"
	"net/mail"
)

var ErrorInvalidUserName = fmt.Errorf("error invalid user name")
var ErrorInvalidEmail = fmt.Errorf("error invalid email")

type UserInput struct {
	Name           string
	UserName       string
	Email          string
	HashedPassword string
}

func (u UserInput) Valid() error {
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return err
	}

	return nil
}

type User struct {
	Id             string
	Name           string
	UserName       string
	Email          string
	HashedPassword string
}

type Store interface {
	AddUser(user UserInput) (*User, error)
	GetUserByEmailOrUsername(emailOrUsername string) (*User, error)
	GetUserById(id string) (*User, error)
	GetUsers() ([]User, error)
}

type PasswordManager interface {
	HashPassword(password string) (string, error)
	CompareHashAndPassword(hashedPassword, password string) error
}
