package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

// func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello World! -server "))
// 	switch r.Method {
// 	case http.MethodGet:
// 		switch r.URL.Path {
// 		case "/":
// 			w.Write([]byte("This is the index page."))
// 		case "/users":
// 			w.Write([]byte("This is the users page."))

// 		}
// 	default:
// 		w.Write([]byte("404 error unknown"))
// 	}

// }

func main() {
	api := &api{addr: ":8080"}
	
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr: api.addr,
		Handler: mux,
	}
	


	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)
	
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
