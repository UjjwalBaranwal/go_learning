package main

import "fmt"

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)
	go func() {
		chan1 <- "Hello"
	}()
	go func() {
		chan2 <- "World"
	}()

	select {
	case var1 := <-chan1:
		{
			fmt.Println("Value from chan1", var1)
		}
	case var2 := <-chan2:
		{
			fmt.Println("Value from chan 2 ", var2)
		}
	}
}
