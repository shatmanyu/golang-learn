package main

import (
	"fmt"
	"time"
)

// const age:=41 -> // This line will cause an error because constants
// cannot be declared with the `:=` syntax outside a function
const age = 41 // This is the correct way to declare a constant
func main() {
	// constants-> // constants are immutable values
	// they are defined using the `const` keyword
	const pi = 3.14
	const greeting = "Hello, World!"
	// pi := 3.14 // This will cause an error because pi is a constant
	// pi = 3.14159 // This will also cause an error because constants cannot

	fmt.Println("Value of pi:", pi)
	fmt.Println("Greeting:", greeting)
	fmt.Println("Age constant:", age)

	// constants grouping
	const (
		day   = 24 * time.Hour
		month = 30 * day
		year  = 365 * day
	)
	// day := 35 // This will cause an error because day is already defined as a constant
	fmt.Println("One day in seconds:", day.Seconds())
}
