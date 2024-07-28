package main

import "fmt"

func main() {

	customer1 := Customer{id: 1, name: "John", age: 30, address: Address{district: "Türkiye", city: "İstanbul"}}
	customer1.ToString()
	customer1.changeName("Change Name Kullanıldı ")
	customer1.ToString()
	// changeName(&customer1)
	// customer1.ToString()
	/*
	 // customer1 = Customer{id: 1, name: "John", age: 30, address: Address{district: "Türkiye", city: "İstanbul"}}
	 var customer2 = Customer{id: 2, name: "Doe", age: 40}

	 fmt.Println(customer1)
	 fmt.Println(customer2)

	 customer1.age = 31
	 fmt.Println(customer1)
	*/

}

// pass by reference kullanılarak name değiştirildi
// func changeName(customer *Customer) {
// 	customer.name = "Chagne Name Kullanıldı"
// }

func (customer *Customer) changeName(newName string) {
	customer.name = newName
}

type Customer struct {
	id      int64
	name    string
	age     int
	address Address
}

type Address struct {
	district string
	city     string
}

// ToString metodu Customer struct'ı için yazıldı pointer kullanılmadı çünkü printf gibi bir değer kullanarak bir şey değiştirmiyoruz

func (customer Customer) ToString() {
	fmt.Printf("Name : %s , Age : %d \n", customer.name, customer.age)
}
