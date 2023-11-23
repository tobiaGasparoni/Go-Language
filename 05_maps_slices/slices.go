package main

import (
	"log"
	"sort"
)

func main() {
	var mySlice []int

	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	log.Println("Before sorting", mySlice)

	sort.Ints(mySlice)

	log.Println("After sorting", mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers[0:2])
	log.Println(numbers[6:9])
}
