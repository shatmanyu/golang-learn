package main

import "fmt"

func main() {

	age := 40

	switch {
	case age < 18:
		fmt.Println("You are a minor.")
	case age >= 18 && age < 65:
		fmt.Println("You are an adult.")
	case age >= 65:
		fmt.Println("You are a senior citizen.")
	default:
		fmt.Println("Age is not defined.")
	}

	i := 5
	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	default:
		fmt.Println("i is neither 1 nor 2")
	}
	// multiple cases can be handled together
	switch i {
	case 3, 4, 5:
		fmt.Println("i is 3, 4, or 5")
	default:
		fmt.Println("i is not 3, 4, or 5")
	}

	// switch with type assertion
	x := func(i interface{}) {
		// x can hold any type
		// here interface{} is used to hold any type of value
		// we can use type assertion to determine the type of x

		switch v := i.(type) {
		case int:
			fmt.Println("x is an int:", v)
		case string:
			fmt.Println("x is a string:", v)
		case bool:
			fmt.Println("x is a bool:", v)
		default:
			fmt.Println("x is of a different type")
		}

	}
	x(42)      // calling with an int
	x("hello") // calling with a string
}
