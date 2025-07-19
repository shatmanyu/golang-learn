package main

import "fmt"

func main() {
	// if-else example
	age := 40
	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age >= 18 && age < 65 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior citizen.")
	}

	// we can also use a short variable declaration in the if statement
	if name := "John"; name == "John" {
		fmt.Println("Hello, John!")
	} else {
		fmt.Println("You are not John.")
	}
	// ternary-like operation using if-else not supported in Go
	// but we can achieve similar functionality with if-else
}
