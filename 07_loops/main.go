package main

import "log"

func main() {
	for i := 0; i <= 5; i++ {
		log.Println(i)
	}

	letters := []string{"a", "b", "c", "d"}

	for i, letter := range letters {
		log.Println(i, letter)
	}

	for _, letter := range letters {
		log.Println(letter)
	}

	myMap := make(map[string]string)

	myMap["dog"] = "Fido"
	myMap["cat"] = "Fluffy"

	for animalType, animal := range myMap {
		log.Println(animalType, animal)
	}

	var text = "Tobia G."
	for i, l := range text {
		log.Println(i, ":", l)
	}
}
