package main

import (
	"log"
	"time"
)

var s = "seven"

var firstName string
var lastName string
var phoneNumber string
var age int
var birthDate time.Time

func main() {
	var s2 = "six"

	log.Println("s is", s)
	log.Println("s2 is", s2)

	saySomething("xxx")

	firstName = "Tobia"
	lastName = "Gasparoni"
	phoneNumber = "+57 123-132-5431"
	age = 24
	birthDate = time.Date(1999, time.May, 6, 13, 30, 0, 0, time.UTC)
	createdAt := time.Now()

	log.Println("Data created at", createdAt)
	log.Println("My birthDate is", birthDate)
}

func saySomething(s3 string) (string, string) {
	log.Println("s from the saySomething func is", s)
	return s3, "World"
}
