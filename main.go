package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{}

func createUserHandler(w http.ResponseWriter, r *http.Request){
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	user.ID = fmt.Sprintf("%d", len(users) + 1)
	users = append(users, user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}


func getUsersHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	id := params["id"]

	for i, user := range users{
		if user.ID == id{
			json.NewDecoder(r.Body).Decode(&users[i])
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}
	http.Error(w, "Пользователь не найден", http.StatusNotFound)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id := params["id"]

	for i, user := range users{
		if id == user.ID{
			users = append(users[:i], users[i + 1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Пользователь не найден", http.StatusNotFound)
}

func main() {

}