package main

import (
	"fmt"
)

func main() {
	/*
		var x int
		fmt.Println(x) // 0 atacaktır.
		var pointer *int
		fmt.Println(pointer) // <nil> atacaktır.

		if pointer == nil {
			fmt.Println("Pointer her hangi bir yeri göstermiyor.")
		}
	*/

	/*
		fileData, err := os.ReadFile("file.txt")
		// if we do not check error fileData , _ := os.ReadFile("file.txt")
		if err != nil {
			// log.Fatal(err)
			// fmt.Println(err)
			// fmt.Println("Dosya okunurken bir hata oluştu.")
			panic(err)
			} else {
				fmt.Println(string(fileData))
			}
	*/

	/*
		result, error := divide(10, 2)
		if error != nil {
			fmt.Println(error)
			} else {
				fmt.Println(result)
			}
	*/

	productService := ProductService{}
	err := productService.Add(Product{id: 1, name: "a", price: -5000})
	if err != nil {
		fmt.Println(err)
	}
}

type Product struct {
	id    int
	name  string
	price int
}
type ProductService struct {
}

func (productService *ProductService) Add(product Product) error {
	if len(product.name) == 0 {
		return ValidationError{code: "name", text: "Ürün ismi boş olamaz."}
	}
	if product.price < 0 {
		return ValidationError{code: "price", text: "Ürün fiyatı 0'dan küçük olamaz."}
	}
	fmt.Println("Ürün eklendi.")
	return nil
}

type ValidationError struct {
	code string
	text string
}

func (validationError ValidationError) Error() string {
	return fmt.Sprintf("Error code: %s, Error text: %s", validationError.code, validationError.text)
}

// func divide(x int, y int) (int, error) {
// 	if y == 0 {
// 		return 0, errors.New("0'a bölme hatası")
// 	}
// 	return x / y, nil
// }
