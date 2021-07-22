package main

import "fmt"

func main() {

	//channel for buffering upto 4 values
	messages := make(chan string, 4)

	messages <- "buffered"
	messages <- "channel"
	messages <- "3-values"
	messages<- "hello"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}