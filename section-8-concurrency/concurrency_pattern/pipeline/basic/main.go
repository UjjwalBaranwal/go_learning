package main

import "fmt"

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}
func doSqr(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
func main() {
	//inpuyt
	input := []int{1, 2, 3, 4, 5}
	// stage 1
	dataChannel := sliceToChannel(input)
	// stage 2
	doSquare := doSqr(dataChannel)
	// stage 3
	for n := range doSquare {
		fmt.Println(n)
	}
}
