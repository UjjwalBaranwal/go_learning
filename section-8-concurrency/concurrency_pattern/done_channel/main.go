package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) { //done channel is read only channel
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing Work")
		}
	}
}

/*
chan bool ----> bidirectional(send + receive)
<-chan bool ---> recieve-only (read channel)
chan <- bool ----> send oly channel(write only)
*/
func main() {
	// for handling goroutine leak
	done := make(chan bool)
	go doWork(done)
	time.Sleep(time.Second * 3)
	close(done)
	//“Goroutines should always have a controlled exit mechanism”
}
