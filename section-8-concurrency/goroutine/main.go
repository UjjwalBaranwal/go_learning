package main

import (
	"fmt"
	"time"
)

func sayHello(msg string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println(msg)
}

func main() { // main() is also a goroutine (the main one)
	fmt.Println("First Line of Main")
	go sayHello("Hello 1", time.Second)
	go sayHello("Hello 2", 2*time.Second)
	go sayHello("Hello 3", 3*time.Second)
	fmt.Println("Last Line")
	time.Sleep(time.Second * 2)
	// Main function exits → all goroutines stop
}
