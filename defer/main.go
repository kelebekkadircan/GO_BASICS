package main

import "fmt"

func main() {

	// fonksiyonun sonunda çalıştırılacak kodu belirler
	// dosya kapanması gerekir bağlantı kapanması gerekir
	//  temizlik amacı ile kullanılır channel açıldıysa kapatılması gerekir
	// defer fmt.Println("Kadircan")
	// fmt.Println("Kelebek")

	defer fmt.Println("1.defer")
	defer fmt.Println("2.defer")
	defer fmt.Println("3.defer")

	fmt.Println("Burası Main Fonksiyonu")
	// deferler stack mantığı ile çalışır
	// en son yazılan en önce çalışır

	defer fmt.Println("4.defer")
	var condition = false
	if condition {
		return
	}
	defer fmt.Println("5.defer")

	x := 10
	y := 20
	// deferler çalıştırılmadan önce değerleri alır
	defer fmt.Println(x, "x")
	defer fmt.Println(y, "y")

	x = 100
	y = 200

	fmt.Println("x : ", x)
	fmt.Println("y : ", y)

	var condition2 = true

	defer cleanUp()
	if condition2 {
		panic("Bir hata oluştu")

	}
	cleanUp()
}

func cleanUp() {
	fmt.Println("CleanUp çalıştı")
}
