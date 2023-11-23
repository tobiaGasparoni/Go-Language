package main

import (
	"log"

	"github.com/tobiaGasparoni/myniceprogram/helpers"
)

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "Some name"
	log.Println(myVar.TypeName)
}
