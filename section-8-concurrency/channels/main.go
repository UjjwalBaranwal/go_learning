package main

import (
	"fmt"
	"time"
)

type User struct {
	name string
}

func main() {

	message := make(chan string) //created a channel of type string
	users := make(chan User)     // unbuffered channel
	go func() {
		fmt.Println("Sending message to channgel")
		message <- "Hello from channel"
	}()

	go func() {
		fmt.Println("Adding in Users channel")
		users <- User{name: "Ujjwal"}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Reading from the channel")
	msg := <-message
	fmt.Println(msg)
	user := <-users
	fmt.Println(user)
}
