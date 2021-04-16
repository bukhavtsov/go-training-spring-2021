package main

import (
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/helloWorld", hello)
	r.Use(mux.CORSMethodMiddleware(r))

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Server doesn't listen port...", err)
	}
	if err := http.Serve(listener, r); err != nil {
		log.Fatal("Server has been crashed...", err)
	}
}

func hello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello World"))
}
