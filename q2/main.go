package main

import (
	"fmt"
)

// channel
func main() {
	/*
	   	ch := make(chan int)

	   	//send a value to the channel
	   	go send(ch)

	   	//receive a value from the channel
	   	x := <-ch
	   	fmt.Println(x)
	   }
	   func send(ch chan int) {
	   	fmt.Println("Sending")
	   	ch <- 90
	   	close(ch)
	   	fmt.Println("sent")
	*/

	//Buffered channel
	ch := make(chan int, 5)

	go send(ch)

	for i := range ch {
		fmt.Printf("received %d from channel\n", i)
	}

}
func send(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("sent %d to channel", i)
	}
	close(ch)
}
