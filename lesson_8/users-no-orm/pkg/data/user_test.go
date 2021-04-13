package data

import (
	"database/sql"
	"errors"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	return db, mock
}
var testUser = &User {
	Id: 1,
	Name: "ExampleName",
	Surname: "ExampleSurname",
}

func TestReadAll(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewUserData(db)
	rows := sqlmock.NewRows([]string{"id", "name", "surname"}).AddRow(testUser.Id, testUser.Name, testUser.Surname)
	mock.ExpectQuery(readAllUsersQuery).WillReturnRows(rows)

	users, err := data.ReadAll()
	assert.NoError(err)
	assert.NotEmpty(users)
	assert.Equal(users[0], *testUser)
	assert.Len(users, 1)
}

func TestReadAllErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewUserData(db)
	mock.ExpectQuery(readAllUsersQuery).WillReturnError(errors.New("something went wrong..."))
	users, err := data.ReadAll()
	assert.Error(err)
	assert.Empty(users)
}
