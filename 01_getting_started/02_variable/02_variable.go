package main

import "fmt"

func main () {
	var name string = "Golang"

	var typeBoolean bool
	typeBoolean = true

	number := 25

	dynamicName := "Doohan" 
	dynamicName = "Aditya"

	fmt.Println(name) // expected: "Golang"
	fmt.Println(typeBoolean) // expected: true
	fmt.Println(number) // expected: 25
	fmt.Println(dynamicName) // expected: "Aditya"
}
