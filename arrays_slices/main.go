package main

import "fmt"

func main() {
	/*
		var names [4]string
		names = [4]string{"John", "Paul", "George", "Ringo"}

		// names[0] = "John"
		// names[1] = "Paul"
		// names[2] = "George"

		fmt.Println(names[0:2])
	*/

	var names = []string{"Kadircan", "Kelebekcan", "George", "Ringo"}

	// slice değişkenine yeni bir eleman eklemek için append fonksiyonu kullanılır.
	names = append(names, "Paul")
	fmt.Println(names)

}
