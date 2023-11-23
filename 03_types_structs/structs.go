package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
	CreatedAt   time.Time
}

func main() {
	user := User{
		FirstName:   "Tobia",
		LastName:    "Gasparoni",
		PhoneNumber: "+57 123-132-5431",
		Age:         24,
		BirthDate:   time.Date(1999, time.May, 6, 13, 30, 0, 0, time.UTC),
		CreatedAt:   time.Now(),
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate, "createdAt:", user.CreatedAt)
}
