package main

import "fmt"

func main() {
	// for loop example
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // skip the rest of the loop when i is 2
		}
		if i == 4 {
			break // exit the loop when i is 4
		}
		fmt.Println("Current value of i:", i)
	}
	// while loop example (using for)
	j := 0
	for j < 5 {
		fmt.Println("Current value of j:", j)
		j++ // increment j
	}
	// nested for loop example
	arr := []string{"apple", "banana", "cherry"}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i == j {
				continue // skip when i equals j
			}
			fmt.Println("Comparing", arr[i], "and", arr[j])
		}
	}

	// infinite loop example
	// Uncomment the following lines to see an infinite loop in action
	// for {

	// }
	// This will run indefinitely, so be careful when using it

	// continue and break examples in for loop
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue // skip the rest of the loop when i is 3
		}
		fmt.Println("Current value of i in continue/break example:", i)
		if i == 4 {
			break // exit the loop when i is 4
		}
	}

	// using range in for loop
	for index, value := range arr {
		fmt.Println("index and value-->", index, value)
	}

	for index := range 5 {
		fmt.Println("value====", index)
	}

}
