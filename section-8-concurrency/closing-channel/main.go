package main

import (
	"fmt"
	"sync"
)

func main() {
	// jobs := make(chan int, 5)
	// done := make(chan bool)

	// go func() {
	// 	for {
	// 		r, ok := <-jobs
	// 		if ok {
	// 			fmt.Println("Recieved Message : ", r)
	// 		} else {
	// 			fmt.Println("Channel closed")
	// 			done <- true
	// 			return
	// 		}
	// 	}
	// }()
	// for i := 1; i <= 3; i++ {
	// 	jobs <- i
	// 	fmt.Println("Sending msg ", i)
	// }
	// close(jobs)
	// fmt.Println(<-done)
	// -- using waitGroup for the same thing

	jobs := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			r, ok := <-jobs
			if ok {
				fmt.Println("Recieved Message : ", r)
			} else {
				fmt.Println("Channel closed")
				return
			}
		}
	}(&wg)
	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("Sending msg ", i)
	}
	close(jobs)
	wg.Wait()
}
