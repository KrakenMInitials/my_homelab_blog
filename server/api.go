package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type api struct {
	addr string
}

var users = []User{}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Recieved request to get Users:\n"))
	
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(users)
	if (err != nil){
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// <a *api> isn't required 
func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a user!!"))

	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := User{
		FirstName: payload.FirstName,
		LastName: payload.LastName,
	}

	err = insertUser(u)
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func insertUser(u User) error {
	if u.FirstName == "" {
		return errors.New("First name is required.")
	}
	if u.LastName == "" {
		return errors.New("Last name is required.")
	}

	for _, user := range users {	
		if user.FirstName == u.FirstName && user.LastName == u.LastName{
			return errors.New("User already exists")
		}
	}

	users = append(users, u)

	return nil
	
}


