package main

import "fmt"

func main() {
	/*
		// Declaring a pointer
		var p *int

		// Declaring a variable
		var a int = 10

		// Assigning the address of a to p
		p = &a
		println(a) // 10
		*p = 20

		// Dereferencing the pointer
		println(&a)
		println(*p)
		println(p)
		println(a)
		println(&p)
	*/

	/*
		a := 10
		var b int
		var p *int
		b = a
		p = &a

		*p = 20
		println(a, b)
	*/

	/*

		a := 10
		println(a)
		println(add12(a))

		println(a)
		add12pointer(&a)
		println(a)
	*/

	var numbers = []int{1, 2, 3}
	fmt.Println(numbers)
	changeValue(numbers)
	fmt.Println(numbers)

}

func changeValue(numbers []int) {
	numbers[0] = 100
}

//// pass by value
// func add12(x int) int {
// 	x = x + 12

// }
// // pass by reference
// func add12pointer(x *int) {
// 	*x = *x + 12
// }
