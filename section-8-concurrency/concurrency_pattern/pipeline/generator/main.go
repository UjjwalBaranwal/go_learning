package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func repeatFunc[T any, K int](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():

			}
		}
	}()
	return stream
}
func take[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	taken := make(chan T)
	go func() {
		defer close(taken)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream:

			}
		}
	}()
	return taken
}
func primeFinder(done <-chan int, randomIntStream <-chan int) <-chan int {
	isPrime := func(randInt int) bool {
		for i := randInt - 1; i > 1; i-- {
			if randInt%i == 0 {
				return false
			}
		}
		return true
	}
	primes := make(chan int)
	go func() {
		defer close(primes)
		for {
			select {
			case <-done:
				return
			case randomInt := <-randomIntStream:
				if isPrime(randomInt) {
					primes <- randomInt
				}
			}
		}
	}()
	return primes
}

func fanIn[T any](done <-chan int, channel ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	fannedInStr := make(chan T)

	transfer := func(c <-chan T) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case fannedInStr <- i:
			}
		}
	}
	for _, c := range channel {
		wg.Add(1)
		go transfer(c)
	}
	go func() {
		wg.Wait()
		close(fannedInStr)
	}()
	return fannedInStr
}
func main() {
	start := time.Now()
	done := make(chan int)
	defer close(done)

	randNumFetcher := func() int {
		return rand.Intn(50000000)
	}
	repeatFunc(done, randNumFetcher)
	// for rando := range repeatFunc(done, randNumFetcher) {
	// 	fmt.Println(rando)
	// }
	randomNumStream := repeatFunc(done, randNumFetcher)
	// naive approach -> not having capability of fanning out
	// primeIntStream := primeFinder(done, randomNumStream)
	// for rando := range take(done, primeIntStream, 10) {
	// 	fmt.Println(rando)
	// }

	// Fan Out
	CPUCOUNT := runtime.NumCPU()
	fmt.Println("Current CPU count : ", CPUCOUNT)
	primeFinderChannel := make([]<-chan int, CPUCOUNT)
	for i := range CPUCOUNT {
		primeFinderChannel[i] = primeFinder(done, randomNumStream)
	}

	// Fan In
	fannedInStream := fanIn(done, primeFinderChannel...)
	for i := range take(done, fannedInStream, 10) {
		fmt.Println(i)
	}
	fmt.Println("Time Taken", time.Since(start))
}
