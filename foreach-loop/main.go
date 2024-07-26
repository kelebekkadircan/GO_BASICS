package main

import "fmt"

func main() {

	//var numbers = []int{1, 2, 3, 4, 5}

	//// classic for loop
	// for index := 0; index < len(numbers); index++ {
	// 	println(numbers[index])
	// }

	// // _ is used to ignore the index  value
	// for _, value := range numbers {
	// 	println(value)
	// }

	/*
		var language = "GoLang"

		for _, character := range language {
			fmt.Printf("Character %c\n", character)
		}
	*/

	names := map[string]int{
		"Kadir":   10,
		"Can":     20,
		"Kelebek": 30,
	}

	for key, value := range names {
		fmt.Printf("Key : %s, Value : %d\n", key, value)
	}

}
