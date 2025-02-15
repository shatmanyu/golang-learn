package main

import (
	"fmt"
	"math"
	"strings"
)

var someName = "fgege"

func sayGreeting(s string) {
	fmt.Println("indsy greeting", s)
}
func function1(arr []string, f func(string)) {
	for _, v := range arr {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func getInitials(s string) (string, string) {
	a := strings.Split(s, " ")
	return a[0], a[1]
}
func main() {
	// var st string = "shattu"
	// var ste = 'f'
	// var st1 = "shattu1"
	// var st2 string
	// fmt.Println(st)
	// fmt.Println(ste)

	name := "Fee"

	var age int = 20
	// var age1 = 30
	// age2 := 40
	// fmt.Println("Hello, Shatmanyu", st1, st2, name)
	// fmt.Println("Hello, Shatmanyu", age, age1, age2)

	// // bits & memory
	// var numOne int = 214
	// var numTwo = -11234
	// numThree := 34
	// fmt.Println("Hello, Shatmanyu", numOne, numTwo, numThree)

	// var score float32 = 25.78
	// var score1 float32 = 2356.56
	// score2 := 232435.56

	// fmt.Println("Hello, Shatmanyu", score, score1, score2)

	fmt.Print("hello, \n")
	fmt.Print("world!")
	fmt.Println("hello ninjas")
	fmt.Println("world ninjas!", age, "and my name is", name)

	// printf formatted strings  %_ = format specifier

	fmt.Printf("mytttt age is %v andd my names is %v", age, name)

	//sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("tetetette", str)

	// var ages [3]int = [3]int{20,25,30}

	var ages1 = [3]int{21, 26, 31}

	var names = [3]string{"21", "26", "ghrgg"}

	fmt.Println(ages1, names, len(ages1), len(names))

	// slices(use array under the hood)

	arr := []string{"1sha", "2gerg", "3ft", "4rert"}

	for i := 0; i < len(arr); i++ {
		if i == 1 {
			fmt.Println("inside if in for loop")
			break
		}
		fmt.Println("iterable is:  ", arr[i])
	}

	fmt.Println(age <= 20)
	fmt.Println(age > 20)

	if age < 20 {
		fmt.Println("age is less than 20")
	} else if age == 20 {
		fmt.Println("age is 20")
	} else {
		fmt.Println("age is greater than 20")
	}

	// for index, value := range arr {
	// 	fmt.Println("element at :", index, "is ", value)
	// }

	// for _, value := range arr {
	// 	value = "new strn"
	// 	fmt.Println("element :", value)
	// }
	// fmt.Println(arr)

	sayGreeting("shattu")
	sayGreeting("pranav")

	function1(arr, sayGreeting)
	fmt.Println("area of circle: ", circleArea(32.45))
	fmt.Println("area of circle: ", circleArea(4.56))

	v1, v2 := getInitials("Shattu Gupta")
	// fmt.Println(getInitials("shart rt"))
	fmt.Println(v1, v2)
	sayHello("heelo")

	fmt.Println(menu, "-------", menu["pie"])

	// looping maps

	for key, val := range menu {
		fmt.Println("key is--", key, "with value--", val)
	}

	// ints as key type
	pb := map[int]string{
		12: "mario",
		13: "pelagio",
		14: "gialo",
	}
	for key, val := range pb {
		fmt.Println("key is--", key, "with value--", val)
	}
	pb[13] = "twealkio"
	fmt.Println(pb)

}
