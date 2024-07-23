package main

import "fmt"

func main() {
	/*
		// değişken tanımlama
		var productName string
		var quantity int
		var discount float32
		var isInStock bool


		// değişkenlere değer atama
		productName = "Go Language Programming Essentials"
		quantity = 100
		discount = 0.38
		isInStock = true

		// değişkenin tipini öğrenmek için reflect paketini kullanabiliriz
		// değişkenlerin değerlerini ekrana yazdıralım
		fmt.Println(productName, reflect.TypeOf(productName))
		fmt.Println(quantity, reflect.TypeOf(quantity))
		fmt.Println(discount, reflect.TypeOf(discount))
		fmt.Println(isInStock, reflect.TypeOf(isInStock))
	*/

	/*
		// değişken tanımlama ve değer atama : koyarak tanımlama yapabiliriz
		productName := "Go Language Programming Essentials"

		fmt.Println(productName)
	*/
	// mutlaka değişkenin tiğini vermemiz gerekiyorsa aşağıdaki gibi tanımlama yapabiliriz
	// var quantity int64 = 100

	/*

	 */

	// değişken tanımlama
	var productName string
	var quantity int
	var discount float32
	var isInStock bool

	// değişkenlere değer atama
	productName = "Go Language Programming Essentials"
	quantity = 100
	discount = 0.38
	isInStock = true

	fmt.Printf("Product Name: %s\n", productName) // %s string ifade için kullanılır
	fmt.Printf("Quantity: %d\n", quantity)        // %d integer ifade için kullanılır
	fmt.Printf("Discount: %.2f\n", discount)      // %.2f ondalık sayı için kullanılır
	fmt.Printf("Is in Stock: %t\n", isInStock)    // %t boolean ifade için kullanılır

}
