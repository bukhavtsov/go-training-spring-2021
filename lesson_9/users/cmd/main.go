package main

import (
	"github.com/bukhavtsov/go-training-spring-2021/lesson_9/users/pkg/api"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/bukhavtsov/go-training-spring-2021/lesson_9/users/pkg/data"
	"github.com/bukhavtsov/go-training-spring-2021/lesson_9/users/pkg/db"

	"github.com/gorilla/mux"
)

var (
	serverPort     = os.Getenv("SERVER_PORT")

	host     = os.Getenv("DB_USERS_HOST")
	port     = os.Getenv("DB_USERS_PORT")
	user     = os.Getenv("DB_USERS_USER")
	dbname   = os.Getenv("DB_USERS_DBNAME")
	password = os.Getenv("DB_USERS_PASSWORD")
	sslmode  = os.Getenv("DB_USERS_SSL")
)

func init() {
	if serverPort == "" {
		serverPort = ":8080"
	}
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
	// 1. set up connection to database
	conn, err := db.GetConnection(host, port, user, dbname, password, sslmode)
	if err != nil {
		log.Fatalf("can't connect to database, error: %v", err)
	}
	// 2. create router that allows to set routes
	r := mux.NewRouter()
	// 3. connect to data layer
	userData := data.NewUserData(conn)
	// 4. send data layer to api layer
	api.ServeUserResource(r, *userData)
	// 5. cors for making requests from any domain
	r.Use(mux.CORSMethodMiddleware(r))

	// 6. start server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Server Listen port...")
	}
	if err := http.Serve(listener, r); err != nil {
		log.Fatal("Server has been crashed...")
	}
}
