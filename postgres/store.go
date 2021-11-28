package postgres

import (
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/gislihr/userservice"
	"github.com/jmoiron/sqlx"
)

const userTableName = "userservice.user"

var userColumns = []string{"id", "name", "username", "email", "hashed_password"}

type Store struct {
	db *sqlx.DB
}

type dbUser struct {
	Id             string `db:"id"`
	Name           string `db:"name"`
	UserName       string `db:"username"`
	Email          string `db:"email"`
	HashedPassword string `db:"hashed_password"`
}

func (u dbUser) toServiceUser() *userservice.User {
	return &userservice.User{
		Id:             u.Id,
		Name:           u.Name,
		UserName:       u.UserName,
		Email:          u.Email,
		HashedPassword: u.HashedPassword,
	}
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}

func (s *Store) AddUser(user userservice.UserInput) (*userservice.User, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query, args, err := psql.Insert(userTableName).
		Columns("name", "username", "email", "hashed_password").
		Values(user.Name, user.UserName, user.Email, user.HashedPassword).
		Suffix("returning id").
		ToSql()

	if err != nil {
		return nil, err
	}

	var id string
	err = s.db.Get(&id, query, args...)

	if err != nil {
		if strings.Contains(err.Error(), "user_username_key") {
			return nil, userservice.ErrorInvalidUserName
		}

		if strings.Contains(err.Error(), "user_email_key") {
			return nil, userservice.ErrorInvalidEmail
		}
		return nil, err
	}
	return s.GetUserById(id)
}

func (s *Store) GetUserById(id string) (*userservice.User, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query, args, err := psql.Select().
		From(userTableName).
		Columns(userColumns...).
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}

	var res dbUser
	err = s.db.Get(&res, query, args...)

	return res.toServiceUser(), err
}

func (s *Store) GetUserByEmailOrUsername(emailOrUsername string) (*userservice.User, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query, args, err := psql.Select().
		From(userTableName).
		Columns(userColumns...).
		Where(sq.Or{sq.Eq{"username": emailOrUsername}, sq.Eq{"email": emailOrUsername}}).
		ToSql()

	if err != nil {
		return nil, err
	}

	var res dbUser
	err = s.db.Get(&res, query, args...)

	return res.toServiceUser(), err
}

func (s *Store) GetUsers() ([]userservice.User, error) {
	panic("not implemented")
}
