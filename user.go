package userservice

import "net/mail"

type UserInput struct {
	Name     string
	Email    string
	Password string
}

func (u UserInput) Valid() error {
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return err
	}

	return nil
}

type User struct {
	Id    string
	Name  string
	Email string
}

type Store interface {
	AddUser(user UserInput) (*User, error)
	GetUserById(id string) (*User, error)
	GetUsers() ([]User, error)
}
