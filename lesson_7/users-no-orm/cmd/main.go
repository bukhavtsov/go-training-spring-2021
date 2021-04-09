package main

import (
	"github.com/bukhavtsov/go-training-spring-2021/lesson_7/users-no-orm/pkg/data"
	"log"
	"os"

	"github.com/bukhavtsov/museum/back-end/db"
)

var (
	host     = os.Getenv("DB_USERS_HOST")
	port     = os.Getenv("DB_USERS_PORT")
	user     = os.Getenv("DB_USERS_USER")
	dbname   = os.Getenv("DB_USERS_DBNAME")
	password = os.Getenv("DB_USERS_PASSWORD")
	sslmode  = os.Getenv("DB_USERS_SSL")
)

func init() {
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "1001"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "postgres"
	}
	if password == "" {
		password = "postgres"
	}
	if sslmode == "" {
		sslmode = "disable"
	}
}


func main() {
	conn, err := db.GetConnection(host, port, user, dbname, password, sslmode)
	if err != nil {
		log.Fatalf("can't connect to database, error: %v", err)
	}
	userData := data.NewUserData(conn)
	user := data.User{
		Name:    "Jacobe",
		Surname: "Morts",
	}
	err = userData.Add(user)
	if err != nil {
		log.Println(err)
	}
	users, err := userData.ReadAll()
	if err != nil {
		log.Fatalf("got an error when tried to call ReadAll method: %v", err)
	}
	log.Println(users)
}
