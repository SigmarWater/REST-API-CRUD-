package main

import (
	"encoding/json"
	"net/http"
	"fmt"
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

func main() {
	
}