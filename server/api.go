package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type api struct {
	addr string
	db *sql.DB
}

var db *sql.DB

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("Error loading .env file:", err)
	}

	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")

	if host == "" || username == "" || password == "" {
		log.Println("Missing essential environment variables.")
		return
	}
	port := "5432"
	dbName := "postgres"
	dsn := "postgres://" + username + ":" + password + "@" + host + ":" + port + "/" + dbName + "?sslmode=require"

	var err error
	db, err = sql.Open("pgx", dsn)
		if err != nil {
		log.Fatal("Failed to open DB:", err)
	}

	//ping test
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	log.Println("Connected to RDS PostgreSQL")

	//create table
	_, err = db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			first_name TEXT NOT NULL,
			last_name  TEXT NOT NULL,
			UNIQUE (first_name, last_name)
		);
	`)
	if err != nil {
		log.Fatal("Failed to ensure users table:", err)
	}
}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Recieved request to get Users:\n"))

	rows, err := db.Query(`SELECT first_name, last_name FROM users ORDER BY id ASC`)
	if err != nil {
		http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.FirstName, &u.LastName); err != nil {
			http.Error(w, "Scan error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(users)
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

	_, err = db.Exec(`INSERT INTO users (first_name, last_name) VALUES ($1, $2)`, u.FirstName, u.LastName)
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// func insertUser(u User) error {
// 	if u.FirstName == "" {
// 		return errors.New("First name is required.")
// 	}
// 	if u.LastName == "" {
// 		return errors.New("Last name is required.")
// 	}

// 	for _, user := range users {	
// 		if user.FirstName == u.FirstName && user.LastName == u.LastName{
// 			return errors.New("User already exists")
// 		}
// 	}

// 	users = append(users, u)

// 	return nil
	
// }


