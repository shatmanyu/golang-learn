package main

import "fmt"

func main() {
	// This is the entry point of the Go application
	// You can add your code here to start the application
	fmt.Println("Hello, Go!")
	// simple values

	fmt.Println(1 + 2)

	// strings

	fmt.Println("Hello, " + "World!")

	// bool
	fmt.Println(true && false)

	// float
	fmt.Println(3.14 * 2)

	// variables

	var name string = "Alice"
	fmt.Println("Name:", name)

	// variables with type inference
	var name2 = "Bob"
	fmt.Println("Name2:", name2)

	// or using short variable declaration
	name3 := "Charlie"
	fmt.Println("Name3:", name3)
}
