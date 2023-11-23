package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	var whatToSay string
	var i int

	whatToSay = "Goodbye!!!"
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThing := saySomething()
	fmt.Println("The function returned", whatWasSaid, theOtherThing)
}

func saySomething() (string, string) {
	return "Something", "else"
}
