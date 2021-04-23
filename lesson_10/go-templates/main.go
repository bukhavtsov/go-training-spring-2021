package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ViewData struct {
	Title string
	Users []User
}

type User struct {
	Name string
	Age int
}

func main() {

	data := ViewData{
		Title: "List of users",
		Users: []User {
			User{
				Name: "First",
				Age:  42,
			},
			User{
				Name: "Second",
				Age:  22,
			},
			User{
				Name: "Third",
				Age:  25,
			},
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		templ, err := template.ParseFiles("lesson_10/go-templates/index.html")
		if err != nil {
			fmt.Errorf("got an eror when tried to parse file %v", err)
		}
		templ.Execute(w, data)
	})
	fmt.Println("Server is listening....")
	http.ListenAndServe(":8080", nil)
}
