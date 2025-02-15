package main

import "fmt"

var points = []int{10, 20, 30}

func sayHello(s string) {
	fmt.Println("inside hello function", s)
	for idx, val := range points {
		fmt.Println("value at index: ", idx, "is ", val)
	}
}

var menu = map[string]float64{
	"soup":   4.99,
	"pie":    7.99,
	"salad":  6.99,
	"coffee": 3.22,
}
