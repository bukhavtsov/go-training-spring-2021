package api

import (
	"encoding/json"
	"github.com/bukhavtsov/go-training-spring-2021/lesson_9/users/pkg/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type userAPI struct {
	data *data.UserData
}

func ServeUserResource(r *mux.Router, data data.UserData) {
	api := &userAPI{data: &data}
	r.HandleFunc("/users", api.getAllUsers).Methods("GET")
	r.HandleFunc("/users", api.createUser).Methods("POST")
}

func (a userAPI) getAllUsers(writer http.ResponseWriter, request *http.Request) {
	users, err := a.data.ReadAll()
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get users"))
		if err != nil {
			log.Println(err)
		}
	}
	err = json.NewEncoder(writer).Encode(users)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}


func (a userAPI) createUser(writer http.ResponseWriter, request *http.Request) {
	user := new(data.User)
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if user == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.Add(*user)
	if err != nil {
		log.Println("user hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
