package main

import "fmt"

func main() {
	/*
	 names := make(map[string]int, 0)

	 names["Kadir"] = 10
	 names["Can"] = 20
	 names["Kelebek"] = 30

	 fmt.Println(names)
	 fmt.Println(names["Kadir"])
	*/

	names := map[string]int{
		"Kadir":   10,
		"Can":     20,
		"Kelebek": 30,
	}

	fmt.Println(names)
	delete(names, "Kadir")
	fmt.Println(names)
}
