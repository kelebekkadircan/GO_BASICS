package main

func main() {
	/*
		// Create a buffered channel with a capacity of 3
		ch := make(chan int, 3)

		// Send 3 values to the channel
		// ch <- 1
		// ch <- 2
		// ch <- 3

		ch <- 1
		ch <- 2

		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	*/

	/*

		messages := make(chan string, 2)

		go func() {
			data := <-messages
			fmt.Println("First data : ", data)
			data2 := <-messages
			fmt.Println("Second data : ", data2)
		}()
		messages <- "Hello"
		messages <- "World"
		messages <- "Kadircancancan"

		time.Sleep(time.Second * 1)

		fmt.Println("End of main")
	*/

	/*

		bufferedChannel := make(chan int, 4)

		go func() {
			for i := 0; i < 10; i++ {
				bufferedChannel <- i
				fmt.Println("Send data : ", i)
			}
			close(bufferedChannel)
			}()

			time.Sleep(time.Second * 3)
			for data := range bufferedChannel {

			fmt.Println("Received Data : ", data)
			time.Sleep(time.Second * 1)
		}

	*/
}
