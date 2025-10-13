package main

import "time"

type Blog struct {
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Content string `json:"content"`
}