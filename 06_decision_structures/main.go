package main

import (
	"log"
)

func main() {
	var isTrue bool
	isTrue = false

	if isTrue == true {
		log.Println("isTrue is true")
	} else {
		log.Println("isTrue is not true")
	}

	word := "hi"

	if word == "hello" {
		log.Println("Well hello to you sir")
	} else if word == "hi" {
		log.Println("Hi to you too")
	}

	number := 100

	if number > 99 && isTrue {
		log.Println("The number is indeed greater than 99")
	} else if number == 101 || !isTrue {
		log.Println("Now it is correct")
	}

	switchVar := "cat"

	switch switchVar {
	case "cat":
		log.Println("animal is set to cat")
	case "dog":
		log.Println("animal is set to dog")
	case "duck":
		log.Println("animal is set to duck")
	default:
		log.Println("animal is set to something else")
	}
}
