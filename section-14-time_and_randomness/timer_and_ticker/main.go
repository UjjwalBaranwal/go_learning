package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	TickerExampleExitWithCounter()
	TickerExampleExitWithTimer()
	TimerExample()
}

/*
Ticker → runs REPEATEDLY
Behavior:
1. fires every interval
2. runs forever (until stopped)
*/
func TickerExampleExitWithCounter() {
	// way 1 terminatin ticker using counter
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	counter := 0
	for range ticker.C {
		counter++
		fmt.Println("Tick Tick")
		if counter >= 5 {
			fmt.Println("Ticker is now stopping : stopped")
			fmt.Println("Controlled via counter, exiting loop")
			return
		}
	}
}
func TickerExampleExitWithTimer() {
	// way 2 terminatin ticker using timer
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.Println("Tick Tick")
		case <-timer.C:
			fmt.Println("Timer is done")
			fmt.Println("Controlled via timer, exiting loop")
			return

		}
	}
}

/*
Timer → runs ONCE
Behavior:
1. waits for duration
2. fires only once
3. then stops automatically
*/
func TimerExample() {
	// timer := time.NewTimer(5 * time.Second)
	// <-timer.C
	// fmt.Println("after 5 second this is printing")
	// using go routine here
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-timer.C
		fmt.Println("5 seconds are over")
	}()
	fmt.Println("this is happening inside the main function")
	wg.Wait()
	fmt.Println("program exit TaDa")

}
