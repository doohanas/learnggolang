package main

import "fmt"

func main () {
	var name string = "Golang"

	var typeBoolean bool
	typeBoolean = true

	number := 25

	dynamicName := "Doohan" 
	dynamicName = "Aditya"

	// Multiple declaration of data
	var (
		firstName = "Doohan"
		lastName = "Aditya"
	)

	fmt.Println(name) // expected: "Golang"
	fmt.Println(typeBoolean) // expected: true
	fmt.Println(number) // expected: 25
	fmt.Println(dynamicName) // expected: "Aditya"
	fmt.Println(len(name)) // expected: "6" Amount of chars
	fmt.Println(name[0]) // expected: "byte of 'n'' / 71 "
	fmt.Println(firstName) // expected: "Doohan"
	fmt.Println(lastName ) // expected: "Aditya"
}
