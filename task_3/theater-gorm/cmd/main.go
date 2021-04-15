package main

import (
	"go-training-spring-2021/task_3/theater-no-orm/pkg/data"
	"go-training-spring-2021/task_3/theater-no-orm/pkg/db"

	"log"
	"os"
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
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "Theater_db"
	}
	if password == "" {
		password = "5959"
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
	theaterData := data.NewTheaterData(conn)

	err = theaterData.UpdateAccount(data.Account{
		Id:          11,
		FirstName:   "RRRRR",
		LastName:    "tttttt",
		PhoneNumber: "435345435345",
		Email:       "fknvzdfkjnv",
	})
	if err != nil {
		log.Fatalf("got an error when tried to call UpdateAccount method: %v", err)
	}

	//err = theaterData.DeleteEntry(data.Accounts_, 10)
	//if err != nil {
	//	log.Fatalf("got an error when tried to call DeleteAccount method: %v", err)
	//}

	//err = theaterData.AddLocation(data.Location{
	//	AccountId:   2,
	//	Address:     "WholeStreet",
	//	PhoneNumber: "+7754446389",
	//})
	//if err != nil {
	//	log.Fatalf("got an error when tried to call AddHall method: %v", err)
	//}

	//err = theaterData.AddHall(data.Hall{
	//	AccountId:  2,
	//	Name:       "HappyHall",
	//	Capacity:   500,
	//	LocationId: 2,
	//})
	//if err != nil {
	//	log.Fatalf("got an error when tried to call AddHall method: %v", err)
	//}

	//err = theaterData.AddAccount(data.Account{
	//	FirstName:   "Dim",
	//	LastName:    "Ivanov",
	//	PhoneNumber: "+375296574897",
	//	Email:       "dimaivanov@gmail.com",
	//})
	//if err != nil {
	//	log.Fatalf("got an error when tried to call AddAccount method: %v", err)
	//}

	tickets, err := theaterData.ReadAllTickets()
	if err != nil {
		log.Fatalf("got an error when tried to call ReadAllTickets method: %v", err)
	}
	for _, el := range tickets {
		log.Println(el)
	}

	posters, err := theaterData.ReadAllPosters()
	if err != nil {
		log.Fatalf("got an error when tried to call ReadAllPosters method: %v", err)
	}
	for _, el := range posters {
		log.Println(el)
	}

	users, err := theaterData.ReadAllUsers(data.Account{Id: 1})
	if err != nil {
		log.Fatalf("got an error when tried to call ReadAllUsers method: %v", err)
	}
	for _, el := range users {
		log.Println(el)
	}
}
