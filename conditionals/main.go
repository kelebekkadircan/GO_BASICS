package main

import "fmt"

func main() {
	// if else statements

	var age = 17

	if age >= 18 {
		fmt.Println("You are eligible to vote!")
	} else {
		fmt.Println("You are not eligible to vote!")
	}

	a := 10
	b := 20
	c := 30

	if a >= b && a >= c {
		fmt.Println("a is the largest number")
	} else if b >= a && b >= c {
		fmt.Println("b is the largest number")
	} else {
		fmt.Println("c is the largest number")
	}

}
