package store

import (
	"fmt"
	"os"
	"testing"

	"github.com/gislihr/userservice"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/matryer/is"
)

func getTestDatabase() (*sqlx.DB, error) {
	return sqlx.Connect(
		"postgres",
		"host=localhost user=user password=pass database=userservice sslmode=disable",
	)
}

var testStore PostgresStore

func TestMain(m *testing.M) {
	if testStore.db == nil {
		db, err := getTestDatabase()
		if err != nil {
			fmt.Printf("error opening test database, %v\n", err)
			os.Exit(1)
		}

		err = DropDB(db)
		if err != nil {
			fmt.Printf("error cleaning test database, %v\n", err)
			os.Exit(1)
		}

		err = MigrateDB(db)
		if err != nil {
			fmt.Printf("error migrating test database, %v\n", err)
			os.Exit(1)
		}

		testStore.db = db
	}
	code := m.Run()
	os.Exit(code)
}

func TestAddUser(t *testing.T) {
	is := is.New(t)
	user, err := testStore.AddUser(userservice.UserInput{
		Name:           "Gisli Hrafnkelsson",
		UserName:       "gisli",
		Email:          "gisli@something.com",
		HashedPassword: "hashed-passsword",
	})

	is.NoErr(err)
	is.True(user != nil)
	is.Equal("Gisli Hrafnkelsson", user.Name)
}
