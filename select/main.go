package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	var data1, data2 string

	go func() {
		time.Sleep(time.Second * 4)
		channel1 <- "Hello"

	}()
	go func() {
		time.Sleep(time.Second * 2)
		channel2 <- "World"
	}()

	for len(data1) == 0 || len(data2) == 0 {

		select {
		case data1 = <-channel1:
			fmt.Println("Data 1: ", data1)
		case data2 = <-channel2:
			fmt.Println("Data 2: ", data2)
		default:
			fmt.Println("No data")
		}
		time.Sleep(time.Second * 1)
	}

}
