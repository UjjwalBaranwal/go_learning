package main

import "fmt"

func main() {
	message := make(chan string, 2) // two size buffered channel
	fmt.Println("Sending Messages to buffered channel")
	message <- "Msg1"
	message <- "Msg2"
	// message <- "Msg3"
	fmt.Println("Getting msesage from the channel")
	fmt.Println(<-message)
	fmt.Println(<-message)
	// fmt.Println(<-message)
}
