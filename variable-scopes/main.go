package main

import "fmt"

var x = 10 // global scope

func main() {
	/*
		condition := true
		if condition {
			// block scope
			x := 10
			fmt.Println(x)
		}
		// fmt.Println(x) // x is not accessible here
	*/

	// condition := true
	// if condition {
	// 	fmt.Println(condition)
	// }

	// fmt.Println(condition)
	fmt.Println(x)
	test()
	fmt.Println(x)

}

func test() {
	// var koymazsak global değişkeni de değiştirir
	// x := 40
	x = 40 // function scope global x değişkenini de değiştirir

	fmt.Println(x)
}
