package main

import (
	"log"
	"net/http"
)

func main() {
	api := &api{addr: ":8080"}
	
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr: api.addr,
		Handler: mux,
	}
	
	// sets up the server
	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	err := srv.ListenAndServe() // starts the server
	if err != nil {
		log.Fatal(err)
	}

}
