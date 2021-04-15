package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetConnection(host, port, user, dbname, password, sslmode string) (*sql.DB, error) {
	args := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, user, dbname, password, sslmode)
	connection, err := sql.Open("postgres", args)
	if err != nil {
		return nil, fmt.Errorf("got an error when tried to make connection with database:%w", err)
	}
	return connection, nil
}
