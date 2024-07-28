package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*

		go f1()
		go f2()
		fmt.Println("End of main")
		time.Sleep(1 * time.Second)
	*/

	/*
		wg := sync.WaitGroup{}
		wg.Add(2)

		go func() {
			defer wg.Done()
			fmt.Println("f1")

			}()
			go func() {
				defer wg.Done()
				fmt.Println("f2")
				}()

				wg.Wait() // blocked
				fmt.Println("End of main")

	*/
	startTime := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println("f1")
		time.Sleep(3 * time.Second)

	}()
	go func() {
		defer wg.Done()
		fmt.Println("f2")
		time.Sleep(3 * time.Second)

	}()
	go func() {
		defer wg.Done()
		fmt.Println("f3")
		time.Sleep(3 * time.Second)

	}()

	wg.Wait()
	fmt.Println("Passed time:", time.Since(startTime))

}

// func f1() {
// 	fmt.Println("f1")
// 	// time.Sleep(3 * time.Second)
// }
// func f2() {
// 	fmt.Println("f2")
// 	// time.Sleep(3 * time.Second)
// }

// Output: End of main f1 f2 (random order)
// f1 and f2 are executed concurrently with main
// main exits before f1 and f2
