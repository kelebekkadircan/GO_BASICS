package main

import (
	"fmt"
	"time"
)

func main() {
	// myChannel := make(chan string)
	// done := make(chan bool)

	/*
		go func() {
			myChannel <- "Message from channel"
			}()

			// blocking call
			message, isOpen := <-myChannel

			println(message, isOpen)
	*/

	// myChannel <- "Message from channel Second"

	/*

		go func() {
			message := <-myChannel
			fmt.Println(message)
			done <- true
			}()
			go func() {
				myChannel <- "Message from channel"
				}()

				// blocking call
				<-done

				fmt.Println("Hello from main end")

	*/

	//*****************************************************// MULTIPLE
	myChannel := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			myChannel <- i
			fmt.Println("Sent ", i)
			time.Sleep(1 * time.Second)
		}
		close(myChannel)
	}()

	for {
		data, isOpn := <-myChannel
		if !isOpn {
			break
		}
		fmt.Println("Received ", data)
	}

}
