package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]string)

	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	myMap2 := make(map[string]int)

	myMap2["First"] = 1
	myMap2["Second"] = 2

	log.Println(myMap2["First"])
	log.Println(myMap2["Second"])

	myMap3 := make(map[string]User)

	myMap3["T"] = User{
		FirstName: "Trevor",
		LastName:  "Philips",
	}

	log.Println(myMap3["T"].FirstName)
	log.Println(myMap3["T"].LastName)
}
