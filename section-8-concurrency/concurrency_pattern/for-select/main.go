package main

import "fmt"

func main() {

	// way 1
	charChannel := make(chan string, 4)
	chars := []string{"a", "b", "c"}
	for _, s := range chars {
		select {
		case charChannel <- s:
		}
	}

	close(charChannel)
	for res := range charChannel {
		fmt.Println(res)
	}

}
