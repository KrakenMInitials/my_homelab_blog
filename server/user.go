package main

type User struct {
	// when json encoders turn struct User into json, it will be with key 'first_name'
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}