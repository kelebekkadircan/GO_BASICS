package main

func main() {

	// total := add(10, 20)
	// fmt.Println(total)

	// addition, subtraction := calculation(10, 20)
	// fmt.Println(addition, subtraction)

	/*
		var numbers = []int{10, 20, 30, 40, 50}

		toplam := sum(numbers)
		fmt.Println(toplam)
	*/

	// sumNumber fonksiyonu içerisine istediğimiz kadar sayıyı gönderebiliriz.
	toplam := sumNumbers(10, 20, 30, 40)
	println(toplam)

}

func sumNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total

}

// func sum(numbers []int) int {
// 	total := 0
// 	for _, number := range numbers {
// 		total += number
// 	}
// 	return total
// }

// func calculation(x int, y int) (int, int) {
// 	return x + y, x - y

// }

// func add(x int, y int) int {
// 	// fmt.Println(x + y)
// 	return x + y

// }
